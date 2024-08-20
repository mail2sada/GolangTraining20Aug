package main

import "fmt"

func main() {
	fmt.Println("Arrays in go")

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(a)
	fmt.Println(len(a))
}
