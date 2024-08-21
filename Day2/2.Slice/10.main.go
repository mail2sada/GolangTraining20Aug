package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice insert")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//[]]int{1, 2, 3, 4,11,5, 6, 7, 8, 9, 10}

	idx := 4

	slicea := slice[:idx+1]
	sliceb := slice[idx:]

	fmt.Println(slicea, sliceb)

	slice = append(slice[:idx+1], slice[idx:]...)
	fmt.Println(slice)
	slice[idx] = 11
	fmt.Println(slice)

}
