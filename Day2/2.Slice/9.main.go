package main

import "fmt"

func main() {
	fmt.Println("Demo: slice delete at index")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	idx := 4
	fmt.Println(slice)
	// slicea := slice[:idx]
	// sliceb := slice[idx+1:]

	// fmt.Println(slicea, sliceb)

	// slice = append(slicea, sliceb...)

	slice = append(slice[:idx], slice[idx+1:]...)

	fmt.Println(slice)

}
