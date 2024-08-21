package main

import "fmt"

func main() {
	fmt.Println("Append inside append")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice = append(append(slice, 1), 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(slice)
}
