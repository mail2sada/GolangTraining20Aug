package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Demo: Function parameters")

	a := Add(100, 200)

	fmt.Println(a)

	fmt.Println(Add(10, 20))

	b := Add(Add(30, 40), Add(50, 60))

	fmt.Println(b)

	fmt.Println(Add(Add(30, 40), Add(150, 160)))

}
