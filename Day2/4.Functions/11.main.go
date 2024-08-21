package main

import "fmt"

func IterateSlice(a []int, fn func(int)) {

	for i, v := range a {
		fmt.Println("Iteration", i)
		fn(v)
	}

}

func main() {

	fmt.Println("Higher order functions")

	sum := 0
	a := func(a int) {
		fmt.Println("Inside annonymous func")
		sum += a
	}

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	IterateSlice(slice, a)

	fmt.Println(sum)
}
