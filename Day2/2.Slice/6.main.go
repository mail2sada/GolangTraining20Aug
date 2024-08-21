package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice")

	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice {
		fmt.Println("i", i, "v", v)
	}

	slice[0] = 100

	slice[1] = 200

	slice[10] = 300

}
