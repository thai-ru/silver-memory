package main

import "fmt"

func main() {
	// Slice literals in Golang
	var x = []int{18, 29, 38}
	fmt.Println(x)

	// reading and writing to a slice
	x[0] = 10
	fmt.Println(x[0])

	//Slice built-in functions

	// len returns the length of the slice
	fmt.Println(len(x))

	// append adds elements to the end of the slice
	x = append(x, 45)
	fmt.Println(x)

	// you can add more than one element at a time
	x = append(x, 46, 87)
	fmt.Println(x)

	// one slice can contain another slice
	y := []int{1, 2, 3, 4, 5}
	x = append(x, y...)
	fmt.Println(x)

	// cap is used to determine the maximum number of elements that can be stored in the slice
	fmt.Println(y, len(y), cap(y))
	y = append(y, 6)
	fmt.Println(y, len(y), cap(y))

	y = append(y, 7, 8, 9)
	fmt.Println(y, len(y), cap(y))

	y = append(y, 10, 36, 8, 39)
	fmt.Println(y, len(y), cap(y))

	f := []int{1, 2, 3, 4}
	g := f[:2]
	h := f[1:]
	f[1] = 20
	g[0] = 10
	h[1] = 30
	fmt.Println("f:", f)
	fmt.Println("g:", g)
	fmt.Println("h", h)

	//	Multi Dimensional Slices
	var mds [][]int
	mds = make([][]int, 10)
	fmt.Println(mds)

	//	looping over a slice
	for i := range mds {
		mds[i] = make([]int, 10)
		fmt.Println(i)
	}

}
