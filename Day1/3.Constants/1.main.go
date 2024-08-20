package main

import "fmt"

func main() {
	fmt.Println("Demo: Constant declaration")
	const c1 int = 100
	fmt.Println(c1)
	const c2, c3, c4 = 100, 200, 300

	const c5, c6, c7 = 10, 10.1, "Hello"
	fmt.Println(c2, c3, c4, c5, c6, c7)

	const (
		a = 100
		b = 3.142
		c = "Constant string"
	)
	fmt.Println(a, b, c)
}
