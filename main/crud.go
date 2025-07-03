package main

import (
	"bufio"
	"fmt"
	"os"
	"simple-registration/fileutils"
	"simple-registration/person"
)

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
