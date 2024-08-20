package main

import "fmt"

func main() {
	fmt.Println("Demo: Logical operators")

	var a, b, c = 10, 20, 30

	var d = a > b && a > c

	fmt.Println(d)

	d = a > b || a < c

	fmt.Println(d)

	d = !(a > b)
	fmt.Println(d)

}
