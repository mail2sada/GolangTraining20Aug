package main

import "fmt"

func main() {
	fmt.Println("Some facts about arrays in go")

	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 3}

	if a1 == a2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	a3 := [3]int{1, 3, 5}

	if a1 == a3 {
		fmt.Println("equal")
	} else {
		fmt.Println("Not equal")
	}

	a4 := [4]int{1, 2, 3, 4}

	if a1 == a4 {
		fmt.Println("")
	}
}
