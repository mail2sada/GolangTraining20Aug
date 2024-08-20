package main

import "fmt"

func main() {
	fmt.Println("Demo iterators")

	a := [...]int{0: 10, 5: 100, 10: 200}

	for i, v := range a {
		v *= i
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(a)
}
