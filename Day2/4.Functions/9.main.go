package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Annonymous functions")

	fn := func() {
		fmt.Println("Hello world!!!")
	}

	fmt.Println(reflect.TypeOf(fn))

	fn()

	fn()

	fn()

}
