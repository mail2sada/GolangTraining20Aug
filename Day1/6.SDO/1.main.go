package main

import "fmt"

func main() {
	fmt.Println("Demo: Short declaration operator :=")

	//var a int = 10

	a := 10
	fmt.Println("Value of a:", a)

	b, c, d := 10, 3.1, "hello"
	fmt.Println(b, c, d)

	fmt.Printf("%T %T %T\n", b, c, d)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
