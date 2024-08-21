package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice len & cap")

	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)

	fmt.Println(len(slice))

	fmt.Println(cap(slice))
}
