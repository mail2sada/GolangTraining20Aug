package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo:Variable declaration")

	var v1, v2, v3 = 10, 1.1, false

	fmt.Println(v1, v2, v3)

	fmt.Println(reflect.TypeOf(v1), reflect.TypeOf(v2), reflect.TypeOf(v3))

	var v4, v5, v6 = -10, "Hello world", true

	fmt.Println(v4, v5, v6)
}
