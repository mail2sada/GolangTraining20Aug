package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice append")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 11)
	fmt.Println("After append 11")
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
