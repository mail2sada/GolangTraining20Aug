package main

import "fmt"

var a = 10

// Exported

func main() {
	fmt.Println("Scope of variable")

	a := 0

	fmt.Println(a)

	if a <= 0 {
		i := 10
		fmt.Println(i)
	}
	//fmt.Println(i)
	{
		x := 10
		a := 100
		fmt.Println(x, a)
	}
	fmt.Println(a)
	demo()
}

func demo() {
	fmt.Println(a)
}
