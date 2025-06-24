package person

import (
	"strconv"
	"testing"
	"time"
)

func TestGetAge(t *testing.T) {
	birthDate, _ := time.Parse(time.DateOnly, "1915-04-07")
	expectedValue := strconv.Itoa(time.Now().Year() - birthDate.Year())
	p := Person{"Bruce Wayne", "Wayne Mansion", "bruceLoves@batmail.com", "(22)9876-5432", birthDate}
	if p.GetAge() != expectedValue {
		t.Errorf(`getAge returned invalid value.  Expected: %s, returned: %s`, expectedValue, p.GetAge())
	}
}
