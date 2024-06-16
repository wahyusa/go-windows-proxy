package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menu() {
	fmt.Println("=====================================")
	fmt.Println("Menu")
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
