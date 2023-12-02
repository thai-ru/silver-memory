package main

import (
	"fmt"
)

func main() {
	// nil map literal
	var nilmap map[string]int
	fmt.Println(nilmap)

	//	empty map literal
	teamWins := map[string]int{}
	fmt.Println(teamWins)

	//	 Non-empty map literal
	teams := map[string][]string{
		"Lions": {"John", "Joseph"},
		"Bulls": {"lona", "Sarah"},
	}
	fmt.Println(teams)

	//	empty map literal
	teamWins["Lions"] = 12
	teamWins["Bulls"]++
	fmt.Println(teamWins["Lions"])
	fmt.Println(teamWins["Bulls"])

	//	ok idiom

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["bye"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	//	Deleting from a map
	delete(m, "hello")
	fmt.Println(m)

	//	Maps as sets
	intSet := map[int]bool{}
	vals := []int{1, 2, 3, 4, 5, 6, 3, 3, 6}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet)
	fmt.Println(len(intSet), len(vals))
	fmt.Println(intSet[5])
	fmt.Println(intSet[7])

	//Prints if it evaluates to true
	if intSet[100] {
		fmt.Println("100 in the set")
	} else {
		fmt.Println("100 not in the set")
	}

}
