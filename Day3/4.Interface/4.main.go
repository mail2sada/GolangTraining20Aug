package main

import "fmt"

func main() {
	fmt.Println("Empty interface...")

	var i interface{}

	var a int

	i = 10
	a1 := i.(int)
	fmt.Println(a1)
	fmt.Println(i)
	i = "Hello"

	n := i.(int)
	fmt.Println(n)

	fmt.Println(i)

	i = 3.14159
	fmt.Println(i)

	i = struct {
		int
		float64
	}{10, 22}
	fmt.Println(i)

}
