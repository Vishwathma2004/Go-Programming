package main

import (
	"fmt"
	"time"
)

type Person struct {
	First string
	Last  string
	Dob   time.Time
}

func NewPerson(first, last string, year, month, day int) *Person {
	return &Person{
		First: first,
		Last:  last,
		Dob:   time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
}

func (p Person) GetAge() int {
	d := time.Since(p.Dob)
	return int(d.Hours() / 24 / 365.25) // Use 365.25 to account for leap years
}

func (p Person) SayHello() {
	fmt.Printf("Hello, %s\n", p.First)
}

func main() {
	p := NewPerson("Jane", "Doe", 1980, 1, 1)
	fmt.Println(p.GetAge()) // Corrected to fmt.Println
	p.SayHello()            // Corrected to call SayHello (capital 'S')
}
