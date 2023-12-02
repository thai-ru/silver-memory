package main

import "fmt"

type Person struct {
	name string
	age  int
	pet  string
}

func main() {
	var fred Person
	fmt.Println(fred)

	julie := Person{
		"Julie",
		40,
		"Cat",
	}

	fmt.Println(julie)

	beth := Person{
		name: "Beth",
		age:  23,
	}

	fmt.Println(beth)

	//Dotted notation for reading and writing into a struct field
	fmt.Println(beth.name)
}
