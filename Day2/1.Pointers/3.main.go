package main

import "fmt"

func main() {
	fmt.Println("Demo pointers")

	var a = 100 // variable

	var ptr *int = &a //pointer

	var ptrptr **int = &ptr //pointer pointer

	fmt.Println(a) // Value of a

	fmt.Println(ptr) // address of a

	fmt.Println(*ptr) //value of a

	fmt.Println(*ptrptr) // address of ptr

	fmt.Println(&ptrptr) //address of ptrptr

	fmt.Println(**ptrptr) //value of a

}
