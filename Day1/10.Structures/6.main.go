package main

import "fmt"

type Annonymous struct {
	int
	float32
	float64
	string
	bool
	a float64
}

func main() {
	fmt.Println("Demo:Annonymous structs fields")

	a := Annonymous{int: 10, float32: 33.1, float64: 44.5, string: "Hello", bool: false}
	fmt.Println(a)
}
