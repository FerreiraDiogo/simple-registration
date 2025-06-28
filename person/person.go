package person

import (
	"fmt"
	"strconv"
	"time"
)

type Person struct {
	Name      string
	Address   string
	Email     string
	Phone     string
	BirthDate time.Time
}

func (p Person) String() {
	fmt.Printf("Name: %s, Age: %s, Address: %s, E-mail: %s, Phone Number: %s\n", p.Name, p.GetAge(), p.Address, p.Email, p.Phone)
}

func (p Person) GetAge() string {
	return strconv.Itoa(time.Now().Year() - p.BirthDate.Year())
}
