package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"simple-registration/fileutils"
	"simple-registration/person"
	"strconv"
	"strings"
	"time"
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

func updatePerson() {
	fmt.Println("Type again the data of the person you want to update. Names are not updatable")
	fileutils.Update(readPersonInput())
}

func removeByName() {
	fmt.Println("Type the name of the person you want to remove")
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, err := sanitize(reader.ReadString('\n'))
	if err != nil {
		panic(err)
	}
	fileutils.Delete(input)
}

func findByName() {

	fmt.Println("Type the name of the person you are looking for. This search is case insentive")

	reader := bufio.NewReader(os.Stdin)
	input, err := sanitize(reader.ReadString('\n'))
	if err != nil {
		panic(err)
	}
	person, err := fileutils.FindByName(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person.String())
	}

}

func listPeople() {
	people := fileutils.List()
	if len(people) == 0 {
		fmt.Println("No one was registered yet.")
	} else {
		fmt.Println("The following people are registered")
		for _, person := range people {
			fmt.Println(person.String())
		}
	}

}

func registerPerson() {
	p := readPersonInput()

	fileutils.Write(readPersonInput())
	fmt.Printf("%s registered with Success!\n\n", p.Name)
}

func readPersonInput() person.Person {
	var name string
	var address string
	var email string
	var phonePrefix string
	var phoneNumber string
	var birthDate string

	running := true
	reader := bufio.NewReader(os.Stdin)

	for running {
		if len(name) == 0 {
			fmt.Print("Insert name:")
			name, _ = sanitize(reader.ReadString('\n'))

			_, nameErr := validateStringInput(name, nil)
			if nameErr != nil {
				fmt.Println(nameErr)
				name = ""
				continue
			}
		}

		if len(address) == 0 {
			fmt.Print("Insert address:")
			address, _ = sanitize(reader.ReadString('\n'))
			_, addErr := validateStringInput(address, nil)
			if addErr != nil {
				fmt.Println(addErr)
				address = ""
				continue

			}

		}
		if len(email) == 0 {
			fmt.Print("Insert email:")
			email, _ = sanitize(reader.ReadString('\n'))
			_, mailErr := validateStringInput(email, validateEmail)
			if mailErr != nil {
				fmt.Println(mailErr)
				email = ""
				continue
			}
		}
		if len(phonePrefix) == 0 || len(phoneNumber) == 0 {
			fmt.Print("Insert phone prefix:")
			phonePrefix, _ = sanitize(reader.ReadString('\n'))
			fmt.Print("Insert phone number:")
			phoneNumber, _ = sanitize(reader.ReadString('\n'))
			_, phoneErr := validateStringInput(phonePrefix+phoneNumber, validatePhoneNumber)
			if phoneErr != nil {
				fmt.Println(phoneErr)
				phonePrefix = ""
				phoneNumber = ""
				continue
			}
		}

		if len(birthDate) == 0 {
			fmt.Print("Insert birthdate:")
			birthDate, _ = sanitize(reader.ReadString('\n'))
			_, ageErr := validateStringInput(birthDate, validateAge)
			if ageErr != nil {
				fmt.Println(ageErr)
				birthDate = ""
				continue
			}
		}

		running = false
	}
	return *person.NewPerson(name, address, email, phonePrefix+phoneNumber, birthDate)
}

// Validates if an input is a valid string and performs aditional
// validation provided by callbacl
func validateStringInput(input string, callback func(string) (bool, error)) (bool, error) {
	if len(input) == 0 || len(strings.TrimSpace(input)) == 0 {
		return false, errors.New("Input Cannot be empty or contain only white Spaces!")
	}
	if callback == nil {
		return true, nil
	}
	return callback(input)
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

func validatePhoneNumber(phoneNumber string) (bool, error) {
	_, parseErr := strconv.ParseInt(phoneNumber, 0, 64)
	if parseErr != nil {
		return false, errors.New("Phone number must contain only numbers")
	}
	if len(phoneNumber) < 10 || len(phoneNumber) > 11 {
		return false, errors.New("Phone number can't have less than 8 or more than 9 digits")
	}
	return true, nil
}

func validateAge(birthDate string) (bool, error) {
	convertedAge, parseErr := time.Parse(time.DateOnly, birthDate)
	if parseErr != nil {
		return false, errors.New("Birthdate format is 'YYYY-MM-DD' ")
	}
	if time.Now().Year()-convertedAge.Year() <= 0 {
		return false, errors.New("Inserted birthdate is invalid. Age Must be greater than 0")
	}
	return true, nil
}

func validateEmail(email string) (bool, error) {
	regex, regexErr := regexp.Compile("^[^@]+@[^@]+\\.[^@]+$")
	if regexErr != nil {
		panic("Regexp error!")
	}
	if !regex.MatchString(email) {
		return false, errors.New("Invalid email. Valid format is 'mail@domain.com'")
	}
	return true, nil
}

func sanitize(input string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}
