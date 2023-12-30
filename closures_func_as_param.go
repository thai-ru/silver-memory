package main

import (
	"fmt"
	"sort"
)

// These are functions inside a functions that can access and modify variable in the outer function.

// Passing functions as parameters to functions
func main() {
	type Animal struct {
		FirstName string
		LastName  string
		Age       int
	}

	animals := []Animal{
		{"John", "Loden", 19},
		{"Jane", "Kamau", 33},
		{"Peter", "Parker", 25},
	}

	fmt.Println(animals)

	sort.Slice(animals, func(i, j int) bool {
		return animals[i].LastName < animals[j].LastName
	})
	fmt.Println(animals)

	sort.Slice(animals, func(i, j int) bool {
		return animals[i].Age < animals[j].Age
	})
	fmt.Println(animals)
}
