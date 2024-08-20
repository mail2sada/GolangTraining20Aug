package main

import "fmt"

func main() {

	fmt.Println("Demo: Swap")

	var a, b = 10, 20

	a, b = b, a

	fmt.Println(a, b)
}
