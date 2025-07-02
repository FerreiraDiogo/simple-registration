package fileutils

import (
	"encoding/json"
	"errors"
	"os"
	"simple-registration/person"
	"strings"
)

const PEOPLE_FILE_PATH = "people.json"

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

		if strings.EqualFold(strings.ReplaceAll(name, " ", ""),
			strings.ReplaceAll(p.Name, " ", "")) {
			return p, nil
		}
	}
	return person.Person{}, errors.New("No person with given name was found")
}
