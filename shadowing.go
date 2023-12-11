package main

import "fmt"

func main() {
	//	Shadowing variables
	x := 10

	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	// Shadowing with multiple assignments

	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)

	fmt.Println(true)
	true := 12
	fmt.Println(true)
}
