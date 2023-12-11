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
}
