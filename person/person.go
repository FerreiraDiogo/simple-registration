package person

import (
	"fmt"
	"strconv"
	"time"
)

type Person struct {
	name      string
	address   string
	email     string
	phone     string
	birthDate time.Time
}

func (p Person) String() {
	fmt.Printf("Name: %s, Age: %s, Address: %s, E-mail: %s, Phone Number: %s\n", p.name, p.GetAge(), p.address, p.email, p.phone)
}

func (p Person) GetAge() string {
	//TODO: cálculo da idade a partir da data de nascimento
	return strconv.Itoa(time.Now().Year() - p.birthDate.Year())
}
