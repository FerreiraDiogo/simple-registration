package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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
