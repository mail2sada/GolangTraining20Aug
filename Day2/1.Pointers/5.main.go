package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Demo: unsafe")

	var a = 100

	ptr := &a

	//	ptrptr := &ptr

	fmt.Println(unsafe.Pointer(&a))

	fmt.Println(ptr)
}
