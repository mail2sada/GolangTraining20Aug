package main

import "fmt"

func main() {
	fmt.Println("Demo:Conditional statements")

	var a, b = 5, 10

	if a > b {
		fmt.Println(a, "is greater than ", b)
	}

	if a < b {
		fmt.Println(a, "is less than ", b)
	}
}
