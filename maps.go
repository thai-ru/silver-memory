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

}
