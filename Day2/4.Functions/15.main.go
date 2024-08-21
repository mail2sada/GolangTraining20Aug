package main

import (
	"fmt"
)

func UpdateSLice(slice []int, a int) {
	slice = append(slice, a)
	fmt.Println(len(slice), cap(slice))
}

func main() {
	fmt.Println("Demo: Slice & functions")

	s := make([]int, 3, 5)

	fmt.Println(len(s), cap(s))
	fmt.Println(s) //[0 0 0]

	UpdateSLice(s, 1)
	fmt.Println(s) //[0,0,0,1]

	UpdateSLice(s, 2)
	fmt.Println(s) //[0,0,0,1,2]

	UpdateSLice(s, 3)
	fmt.Println(s) //[0,0,0,1,2]

}
