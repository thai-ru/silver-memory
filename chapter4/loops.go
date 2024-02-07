package main

import "fmt"

func main() {
	//x := 0
	//for x < 5 {
	//	fmt.Println("Value of x is :", x)
	//	x++
	//}

	usernames := []string{"stacy", "mary", "mercy", "jong", "un"}

	for i := 0; i < len(usernames); i++ {
		fmt.Println("Value of i:", usernames[i])
	}

}
