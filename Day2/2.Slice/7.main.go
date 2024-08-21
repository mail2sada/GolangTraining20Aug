package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice[:5])
	fmt.Println(slice[5:])

	fmt.Println(slice[1:6])
}
