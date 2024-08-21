package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1 // 0, 1 | 1,1| 1, 2|,2,3|,3,5
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	//fmt.Println(f(), f(), f(), f(), f())

	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
