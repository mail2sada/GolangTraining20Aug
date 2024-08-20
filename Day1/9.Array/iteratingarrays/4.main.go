package main

import "fmt"

func main() {
	fmt.Println("Demo iterators")

	a := [...]int{0: 10, 5: 100, 10: 200}

	for i, v := range a {
		fmt.Println(i, v)
	}
}
