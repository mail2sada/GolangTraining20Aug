package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Println("Demo: Unsafe")

	a := 100

	ptr := unsafe.Pointer(&a)

	fmt.Println(ptr)

	var f *float64 = (*float64)(ptr)

	fmt.Println(*f)
}
