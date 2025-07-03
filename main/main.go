package main

import (
	"fmt"
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
	case 2:
		updatePerson()
		printWelcomeMessage()
		return true
	case 3:
		listPeople()
		printWelcomeMessage()
		return true
	case 4:
		findByName()
		printWelcomeMessage()
		return true
	case 5:
		removeByName()
		printWelcomeMessage()
		return true
	default:
		fmt.Println("Invalid Option!")
		printMenu()
		return true
	}
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
