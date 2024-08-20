package main

import "fmt"

func main() {
	fmt.Println("Demo: for loops")

	var a int = 10

	for a >= 0 {
		fmt.Println(a)
		a--
	}
}
