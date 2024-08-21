package main

import "fmt"

func Add(list ...int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Demo variadic arguments")
	fmt.Println(Add(1, 2, 3, 4, 5, 6), Add(1, 2), Add())
	s := Add(Add(1, 2, 3), Add(5, 6, 7))
	fmt.Println(s)
}
