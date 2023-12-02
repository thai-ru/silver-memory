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
		"Lions": []string{"John", "Joseph"},
		"Bulls": []string{"lona", "Sarah"},
	}
	fmt.Println(teams)
}
