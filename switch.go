package main

import "fmt"

func main() {
	//	 Switch statements in Go lang
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "Is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "Is exactly 5 letters long!", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "Is longer than 9 letters")
		}
	}

	//	Blank Switch
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is too short")
		case wordLen > 10:
			fmt.Println(word, "Is a long word!")
		default:
			fmt.Println(word, "Is the right length!")
		}
	}
}
