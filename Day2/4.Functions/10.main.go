package main

import "fmt"

func main() {
	fmt.Println("Closures")

	var data = 0

	a := func() {
		data++
	}

	a()

	a()
	a()

	fmt.Println(data)
}
