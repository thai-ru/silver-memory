package main

import (
	"fmt"
)

// Person --> Declaring a user-defined type with the name Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Method declaration
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {

	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	output := p.String()
	fmt.Println(output)
}
