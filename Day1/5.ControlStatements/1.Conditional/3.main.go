package main

import "fmt"

func main() {
	fmt.Println("Demo: if elseif else")

	var a, b = 10, 100

	if a > b {
		fmt.Println(a, "is greater than ", b)

	} else if a < b {
		fmt.Println(a, "is less than ", b)

	} else {
		fmt.Println(a, "is equal to", b)
	}
}
