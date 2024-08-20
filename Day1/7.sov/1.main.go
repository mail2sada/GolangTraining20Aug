package main

import "fmt"

var a = 100

func main() {
	fmt.Println("Demo scope of variable..")
	i := 0
	fmt.Println(i, a)
}

func demo() {
	fmt.Println(a)
	//fmt.Println(i)
}
