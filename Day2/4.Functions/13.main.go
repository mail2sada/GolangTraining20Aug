package main

import "fmt"

func UpdateArray(a *[3]int) {
	a[0] = 10

	a[1] = 20

	a[2] = 30
}
func main() {

	fmt.Println("Demo: Arrays & functions")

	array := [3]int{1, 2, 3}

	fmt.Println(array)
	UpdateArray(&array)

	fmt.Println(array)
}
