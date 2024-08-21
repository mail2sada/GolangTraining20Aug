package main

import "fmt"

func main() {
	fmt.Println("Demo pointers")

	var array = [3]int{1, 2, 3}

	var ptr *[3]int = &array

	fmt.Println(array)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	fmt.Println(array[0])
	fmt.Println(ptr[0])
	for i, v := range ptr {
		fmt.Println(i, v)
	}
}
