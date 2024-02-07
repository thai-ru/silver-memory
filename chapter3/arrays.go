package main

import "fmt"

func main() {

	// The arrays are zero-initialized
	var x [3]int
	fmt.Println(x)

	// Arrays can be initialized with a literal
	var j = [3]int{23, 6, 9}
	fmt.Println(j)

	// Sparse array --> Array where most elements are set to their zero value
	var k = [7]int{2, 5: 3, 7}
	fmt.Println(k)

	// You can leave of the number of elements in the array
	var g = [...]int{2, 3, 7}
	fmt.Println(g)

	//Comparing arrays using (==) operator and (!=) operator
	fmt.Println(j == g) // Evaluates to false
	fmt.Println(j != g) // Evaluates to true

}
