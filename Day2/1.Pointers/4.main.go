package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	roll_no int
	name    string
}

func main() {
	fmt.Println("Demo: Pointers")

	var s = Student{roll_no: 1, name: "Anil kumar"}

	var ptr *Student = &s

	fmt.Println(ptr)
	fmt.Println(*ptr)

	fmt.Println(unsafe.Pointer(&s))
}
