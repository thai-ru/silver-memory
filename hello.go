package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	//	Strings
	var nameOne string = "John the baptist"
	var nameTwo = "Luigi and mario"
	var nameThree string

	fmt.Println(nameOne)
	fmt.Println(nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "browser"

	fmt.Println(nameOne)
	fmt.Println(nameTwo, nameThree)

	//shorthand can't be used outside a function
	nameFour := "Ninjas"

	fmt.Println(nameFour)

	//	Integer

	var ageOne = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	age := 21
	name := "Robinson"

	fmt.Println("My age is", age)

	//Formatted strings
	fmt.Printf("My name is %v \n", name)

	//	bits and memory

	//var numOne int8 = 25
	//var numTwo int8 = -128
	//var numThree uint8 = 255
	//
	//var scoreOne float32 = 25.86
	//var scoreTwo float64 = 766.099879705

	//fmt.Print("Hello world! \n")
	//fmt.Print("Hello world! \n")
	//fmt.Print("Hello world! \n")

	//	Arrays and slices

	var ages [3]int = [3]int{12, 14, 16}
	//var agesTwo = [3]int{12, 14, 16}

	names := [4]string{"joe", "elias", "job", "abel"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//	Slices (use arrays under the hood)

	scores := []int{100, 200, 30}
	scores[1] = 40

	scores = append(scores, 45)
	fmt.Println(scores)

	//	ranges in slice

	rangeOne := scores[1:2]
	rangeTwo := names[2:]

	rangeTwo = append(rangeTwo, "kim")
	fmt.Println(rangeOne, rangeTwo)

	greeting := "hello there friends"

	fmt.Println(strings.Split(greeting, " "))

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

	users := []int{45, 20, 78, 76, 23, 44, 37, 120}
	sort.Ints(users)
	fmt.Println(users)

	index := sort.SearchInts(users, 140)
	fmt.Println(index)

	usernames := []string{"stacy", "mary", "mercy", "jong", "un"}
	sort.Strings(usernames)
	fmt.Println(usernames)

}
