package main

import "fmt"

func main() {
	fmt.Println("Demo iterators")

	a := [...]int{0: 10, 5: 100, 10: 200}

	for _, v := range a {
		fmt.Println(v)
	}

	for i, _ := range a {
		fmt.Println(i)
	}
}
