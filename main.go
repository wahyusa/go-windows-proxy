package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// isProxyEnabled checks if the proxy is enabled bool
func isProxyEnabled() bool {
	cmd := exec.Command("reg", "query", "HKEY_CURRENT_USER\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	return strings.Contains(string(output), "0x1")
}

func menu() {
	proxyStatus := "Off"
	if isProxyEnabled() {
		proxyStatus = "On"
	}
	fmt.Println("=====================================")
	fmt.Printf("Menu | Status: %s\n", proxyStatus)
	fmt.Println("1. On")
	fmt.Println("2. Off")
	fmt.Println("3. Setting")
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
			// TODO: enableProxy()
			fmt.Println("On")
		case "2":
			// TODO: disableProxy()
			fmt.Println("Off")
		case "3":
			// TODO: setting()
			fmt.Println("Setting")
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
	handleMenu()
}
