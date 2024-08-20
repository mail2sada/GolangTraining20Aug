package main

import "fmt"

func main() {
	fmt.Println("Demo: Assignment operator")

	var a, b = 10, 20

	a += b

	fmt.Println(a)

	a -= b

	fmt.Println(a)
}
