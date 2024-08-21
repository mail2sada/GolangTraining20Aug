package main

import "fmt"

func divide(a, b int) int {
	defer fmt.Println("defer from divide") //3 //6

	f := a / b
	fmt.Println("Returing from divide", f) //2 5
	return f

}

func main() {
	fmt.Println("Demo defer..") //1

	c := divide(10, 2)
	fmt.Println(c) //4
	c = divide(10, 5)
	fmt.Println(c)
	c = divide(10, 0)
	fmt.Println(c)

}
