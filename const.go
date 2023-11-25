package main

import "fmt"

const x int64 = 10

const (
	idkey   = "id"
	namekey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// Go processes constants during compile time
	x = x + 1
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}
