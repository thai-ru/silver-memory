package main

import (
	"fmt"
)

func main() {
	l := 10
	pointerToL := &l
	fmt.Println(l)
	fmt.Println(pointerToL)  // this prints the memory address of l
	fmt.Println(*pointerToL) // dereferencing the memory address of l

	h := 5 + *pointerToL
	fmt.Println(h)
}
