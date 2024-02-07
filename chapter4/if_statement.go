package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("Here is too low:", n)
	} else if n > 5 {
		fmt.Println("Here is too high:", n)
	} else {
		fmt.Println("Good number:", n)
	}
	//fmt.Println("Here n is undefined:", n)
}
