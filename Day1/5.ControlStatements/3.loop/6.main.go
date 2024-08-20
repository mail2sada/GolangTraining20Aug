package main

import "fmt"

func main() {
	fmt.Println("Nested loops")

	var i, j int

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Print(" ", j)
		}
		fmt.Println()
	}
}
