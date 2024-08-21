package main

import "fmt"

func main() {

	fmt.Println("Demo: Pointers...")

	var a int = 100

	var ptr *int = &a

	fmt.Println("Value of a", a)

	fmt.Println("Value of a reference from ptr", *ptr)

	*ptr = 200

	fmt.Println("Value of a is ", a)

	*ptr--
	fmt.Println("Value of a", a)

}
