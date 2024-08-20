package main

import "fmt"

func main() {
	fmt.Println("Demo: if else")

	var a, b = 10, 10

	if a > b {
		fmt.Println(a, " is greater than", b)
	} else {
		fmt.Println(a, "is less than ", b)
	}
}
