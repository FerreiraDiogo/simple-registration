package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {

	running := true
	var menuOption int
	printWelcomeMessage()
	for running {
		_, err := fmt.Scanf("%d", &menuOption)
		if err != nil {
			fmt.Println("Please, type a valid option")
			printMenu()
			continue
		}
		running = selectFeat(menuOption)
	}
	printGoodbyeMessage()
}

func selectFeat(menuOption int) bool {
	switch menuOption {
	case 0:
		return false
	case 1:
		registerPerson()
		printWelcomeMessage()
		return true
	default:
		fmt.Println("Invalid Option!")
		printMenu()
		return true
	}
}

func registerPerson() {
	fmt.Println("Mock a Person registration!")
}

func printGoodbyeMessage() {
	fmt.Println("Thank You for using Simple People Registrator!")
}

func printWelcomeMessage() {
	fmt.Println("==========Simple People Registrator V1.0.0==========")
	printMenu()

}

func printMenu() {
	fmt.Println("========================Menu========================")
	fmt.Printf("1 - Register New Person\n2 - Update People\n3 - List People\n4 - Find People by name\n5 - Remove People by name\n0 - Quit\n")
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // For Windows
	} else {
		cmd = exec.Command("clear") // For Unix-like systems (Linux, macOS)
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
