package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo:Variable declaration")

	var v1 int
	fmt.Println("Value of v1", v1)
	fmt.Println("Type of v1", reflect.TypeOf(v1))

	var v2 = 100
	fmt.Println("Value of variable v2", v2)
	fmt.Println("Type of v2", reflect.TypeOf(v2))

	var v3 = "hello world!!"
	fmt.Println("v3:", v3)
	fmt.Println("Type of v3", reflect.TypeOf(v3))

	var v4 = true
	fmt.Println("v4", v4)
	fmt.Println("Type of v4", reflect.TypeOf(v4))

}
