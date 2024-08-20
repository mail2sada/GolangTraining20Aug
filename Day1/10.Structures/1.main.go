package main

import "fmt"

type FirstStruct struct {
	v1 int
	v2 string
	v3 float32
	v4 bool
}

func main() {
	fmt.Println("Demo structus")

	var a FirstStruct

	fmt.Println(a)
}
