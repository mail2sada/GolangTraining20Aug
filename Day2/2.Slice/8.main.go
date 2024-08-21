package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice")

	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)

	slice = slice[1:]
	fmt.Println(slice)

	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}
