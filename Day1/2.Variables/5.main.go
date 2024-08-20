package main

import "fmt"

func main() {
	fmt.Println("Variable declaration!!")

	var (
		a int     = 100
		b float32 = 2.3
		c string  = "Hello"
		d bool    = false
		e uint    = 500
	)

	fmt.Println(a, b, c, d, e)
}
