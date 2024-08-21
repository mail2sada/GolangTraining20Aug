package main

import "fmt"

func main() {
	fmt.Println("Demo:Slice")

	slice := []int{}

	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		fmt.Println("len", len(slice))
		fmt.Println("cap", cap(slice))
	}
}
