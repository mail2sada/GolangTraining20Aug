package main

import (
	"fmt"
)

func AppendSlice(slice *[]int, a int) {
	//*slice[0] = 1000
	*slice = append(*slice, a)
	//*slice[1] = 5000
}

func main() {

	fmt.Println("Demo slice & function")

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s)
	AppendSlice(&s, 10)
	fmt.Println(s)
}
