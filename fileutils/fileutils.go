package fileutils

import (
	"encoding/json"
	"errors"
	"os"
	"simple-registration/person"
)

const PEOPLE_FILE_PATH = "people.json"

func Write(people person.Person) {

	file, err := open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonified, marshErr := json.MarshalIndent(people, "", " ")
	if marshErr != nil {
		panic(marshErr)
	}
	file.Write(jsonified)
}

func open() (*os.File, error) {
	file, err := os.OpenFile(PEOPLE_FILE_PATH, os.O_CREATE, 0644)
	if err != nil {
		return nil, errors.New("couldn' open file")
	}
	return file, nil
}

func Read() []person.Person {
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
