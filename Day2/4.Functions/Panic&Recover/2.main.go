package main

import "fmt"

func divide(a, b int) int {
	defer func() {
		rec := recover()

		if rec != nil {
			fmt.Println("We go a panic and recovered from here", rec)
		}
	}()

	return a / b
}

func main() {
	fmt.Println("Demo: Panic and recover")

	c := divide(10, 5)
	fmt.Println(c) //2
	c = divide(10, 2)
	fmt.Println(c) //5
	c = divide(10, 4)
	fmt.Println(c) //2
	c = divide(10, 0)
	fmt.Println(c) //0
	c = divide(10, 10)
	fmt.Println(c) //1

}
