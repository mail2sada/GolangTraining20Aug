package main

import "fmt"

func main() {
	fmt.Println("Demo: Multidimentional slice")

	var mslice [][]int = [][]int{{1, 2, 3, 4, 5, 6}, {4, 5, 6, 10, 11, 12, 13, 14}, {7, 8, 9, 1}, {}}

	for i, v := range mslice {
		fmt.Println("i", i, "v", v)
		for j, v1 := range v {
			fmt.Println(j, v1)
		}
	}
}
