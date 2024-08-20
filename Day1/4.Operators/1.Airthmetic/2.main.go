package main

import "fmt"

func main() {
	fmt.Println("Demo: Relational operators")

	var a, b = 10, 5

	var c = a > b

	fmt.Println(c)

	c = a < b
	fmt.Println(c)

	c = a == b
	fmt.Println(c)

	c = a != b
	fmt.Println(c)

}
