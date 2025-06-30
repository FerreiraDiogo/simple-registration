package fileutils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"simple-registration/person"
)

const PEOPLE_FILE_PATH = "people.json"

func Write(person person.Person) {
	// file, err := open()
	// if err != nil {
	// 	panic(err)
	// }
	file, err := os.OpenFile(PEOPLE_FILE_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		errors.New("couldn' open file")
	}
	defer file.Close()
	fmt.Println(person)
	file.WriteString(person.String())
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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return make([]person.Person, 0)
}
