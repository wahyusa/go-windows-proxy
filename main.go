package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

const configFileName = "wproxy.ini"

type ProxyConfig struct {
	Address string
	Enabled bool
}

// loadConfig loads the proxy configuration from the configuration file
func loadConfig() (*ProxyConfig, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		// If the config file doesn't exist, create a new one with default values
		if os.IsNotExist(err) {
			defaultConfig := &ProxyConfig{Address: "26.26.26.26:10809", Enabled: false}
			saveConfig(defaultConfig)
			return defaultConfig, nil
		}
		return nil, err
	}
	defer file.Close()

	config := &ProxyConfig{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "address=") {
			config.Address = strings.TrimPrefix(line, "address=")
		} else if strings.HasPrefix(line, "enabled=") {
			config.Enabled = strings.TrimPrefix(line, "enabled=") == "true"
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

// saveConfig saves the proxy configuration to the configuration file
func saveConfig(config *ProxyConfig) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("address=%s\nenabled=%v\n", config.Address, config.Enabled))
	return err
}

// getConfigFilePath gets the configuration file path
func getConfigFilePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(exePath)
	return filepath.Join(dir, configFileName), nil
}

// isProxyEnabled checks if the proxy is enabled in both registry and config file
func isProxyEnabled() (bool, bool) {
	exactStatus := false
	cmd := exec.Command("reg", "query", "HKEY_CURRENT_USER\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable")

	output, err := cmd.Output()
	if err == nil && strings.Contains(string(output), "0x1") {
		exactStatus = true
	}

	// Load config status
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		// If there's an error loading the config, we can't determine the config status, so return false
		return exactStatus, false
	}

	return exactStatus, config.Enabled
}

func enableProxy() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	config.Enabled = true
	err = saveConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Proxy Enabled")
}

func disableProxy() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	config.Enabled = false
	err = saveConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Proxy Disabled")
}

func setProxyAddress(addr string) {
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	config.Address = addr
	err = saveConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Proxy address set to: %s\n", addr)
}

func addressSettings() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter proxy address: ")
	addr, _ := reader.ReadString('\n')
	addr = strings.TrimSpace(addr)
	setProxyAddress(addr)
}

func menu() {
	clearScreen()

	exactStatus, configStatus := isProxyEnabled()
	exactStatusStr := "Off"
	if exactStatus {
		exactStatusStr = "On"
	}

	configStatusStr := "Off"
	if configStatus {
		configStatusStr = "On"
	}

	fmt.Println("=====================================")
	fmt.Printf("Menu | Exact Status: %s | Config Status: %s\n", exactStatusStr, configStatusStr)
	fmt.Println("1. On")
	fmt.Println("2. Off")
	fmt.Println("3. Address")
	fmt.Println("4. Exit")
	fmt.Println("=====================================")
}

func handleMenu() {
	reader := bufio.NewReader(os.Stdin)
	menu()

	for {
		fmt.Print("Choose: ")
		choose, _ := reader.ReadString('\n')
		choose = strings.TrimSpace(choose)

		switch choose {
		case "1":
			enableProxy()
			menu()
		case "2":
			disableProxy()
			menu()
		case "3":
			addressSettings()
			// After settings, you might want to show the menu again
			handleMenu()
			return // Return after handling to prevent re-execution of the loop
		case "4":
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Wrong input")
		}
	}
}

func main() {
	// Check if command line arguments are provided
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if strings.HasPrefix(arg, "addr=") {
			addr := strings.TrimPrefix(arg, "addr=")
			setProxyAddress(addr)
			// Check for additional arguments for enabling/disabling
			if len(os.Args) > 2 {
				switch os.Args[2] {
				case "on":
					enableProxy()
				case "off":
					disableProxy()
				default:
					fmt.Println("Invalid argument")
				}
			}
			return
		} else {
			switch arg {
			case "on":
				enableProxy()
				return
			case "off":
				disableProxy()
				return
			default:
				fmt.Println("Invalid argument")
				return
			}
		}
	}

	handleMenu()
}
