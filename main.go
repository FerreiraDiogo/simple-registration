package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	running := true
	printWelcomeMessage()
	for running {
		running = false
	}
	clearScreen()
	printGoodbyeMessage()
}

func printGoodbyeMessage() {
	fmt.Println("Thank You for using Simple People Registrator!")
}

func printWelcomeMessage() {
	fmt.Println("==========Simple People Registrator V1.0.0==========")
	fmt.Println("========================Menu========================")
	fmt.Printf("1 - Register New Person\n2 - Update People\n3 - List People\n4 - Find People by name\n5 - Remove People by name\n6 - Quit")

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
