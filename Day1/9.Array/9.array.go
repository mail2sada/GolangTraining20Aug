package main

import "fmt"

func main() {
	fmt.Println("Arrays in go")

	a := [...]int{0: 10, 10: 100, 100: 1000, 1000: 10000, 10000: 100000}
	fmt.Println(len(a))
	fmt.Println(a)
}
