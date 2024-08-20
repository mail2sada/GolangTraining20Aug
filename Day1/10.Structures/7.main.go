package main

import "fmt"

func main() {
	fmt.Println("Annonymous structs")

	a := struct {
		a int
		b float64
		c string
	}{a: 10, b: 3.1, c: "Hello"}

	fmt.Println(a)

	b := a

	b.a = 100
	b.b = 44.0
	b.c = "hi"
	fmt.Println(b)
}
