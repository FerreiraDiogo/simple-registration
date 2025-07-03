package fileutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"simple-registration/person"
	"slices"
	"strings"
)

const PEOPLE_FILE_PATH = "../people.json"

func Write(person person.Person) {

	file, openErr := open()
	if openErr != nil {
		panic(openErr)
	}
	defer file.Close()

	jsonified, marshErr := json.MarshalIndent(person, "", " ")
	if marshErr != nil {
		panic(marshErr)
	}
	file.Write(jsonified)
}

func open() (*os.File, error) {
	file, err := os.OpenFile(PEOPLE_FILE_PATH, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, errors.New("couldn' open file")
	}
	return file, nil
}

func List() []person.Person {
	file, err := open()
	if err != nil {
		panic(err)
	}

	people := make([]person.Person, 0)
	decoder := json.NewDecoder(file)
	for decoder.More() {
		p := person.Person{}
		decoder.Decode(&p)
		people = append(people, p)
	}

	return people
}

func FindByName(name string) (person.Person, error) {

	file, err := open()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var p person.Person
		err := decoder.Decode(&p)
		if err != nil {
			panic(err)
		}

		if namesAreEqual(name, p.Name) {
			return p, nil
		}
	}
	return person.Person{}, errors.New("no person with given name was found")
}

func Delete(name string) {
	people := List()
	var removed []person.Person
	if len(people) == 0 {
		fmt.Println("No one is registered yet")
	} else if len(people) == 1 {
		clearFile()
	} else {
		for index, person := range people {
			if namesAreEqual(name, person.Name) {
				removed = slices.Delete(people, index, index+1)
				break
			}
		}
		if len(removed) > 0 {
			clearFile()
			writeAll(removed)
		}

		fmt.Printf("%s removed successfully\n", name)

	}
}

func clearFile() {
	file, err := open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Truncate(0)

}

func Update(updatedPerson person.Person) {

	people := List()
	for index, person := range people {
		if namesAreEqual(person.Name, updatedPerson.Name) {
			people[index] = updatedPerson
			break
		}
	}
	clearFile()
	writeAll(people)

}

func writeAll(people []person.Person) {
	for _, p := range people {
		Write(p)
	}
}

func namesAreEqual(n1, n2 string) bool {
	return strings.EqualFold(strings.ReplaceAll(n1, " ", ""),
		strings.ReplaceAll(n2, " ", ""))

}
