package main

import "fmt"

func main() {
	fmt.Println("Demo: Arithmetic operators")

	var a, b int = 100, 200

	var c int

	c = a + b
	fmt.Println(c)

	c = a - b
	fmt.Println(c)

	c = a * b
	fmt.Println(c)

	c = a / b

	fmt.Println(c)

	c = a % b
	fmt.Println(c)

}
