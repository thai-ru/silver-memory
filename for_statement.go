package main

import (
	"fmt"
)

func main() {
	//	Go has 4 ways of using for statements
	//	a). Complete for statement

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// b).	Condition only for statement ---> More of a While statement in other languages
	x := 1

	for x < 100 {
		fmt.Println(x)
		x = x * 2
	}

	//	c). Infinite loop
	for {
		fmt.Println("Hello John Doe")
		break
	}

	//	Break, Continue keyword

	for y := 1; y <= 100; y++ {
		if y%3 == 0 && y%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if y%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if y%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(y)
	}
	//	d). for, range loop  --> Only used to loop Go in-built types.(string, arrays, slices and maps)
	evenVals := []int{2, 4, 6, 8, 10, 12}

	//	i --> Index
	//  v --> Value

	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// Just print the value without the index using an underscore (_)
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// Looping through a Set.
	// K --> Key

	uniqueNames := map[string]bool{
		"fred": true,
		"raul": true,
		"aksd": true,
	}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	//	For-Range is a copy
	for _, v := range evenVals {
		// Modifying the value here will not modify the value in the compound type.
		v *= 2
		fmt.Println(v)
	}
	fmt.Println(evenVals)
}
