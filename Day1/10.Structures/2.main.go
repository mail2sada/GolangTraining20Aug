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

	var a FirstStruct = FirstStruct{v1: 10, v2: "hello", v3: 3.1, v4: false}

	fmt.Println(a)
}
