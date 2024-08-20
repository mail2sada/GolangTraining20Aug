package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}

	fmt.Println(a[0], a[1], a[2])
	a[0] = 500

	a[1] = 100
	a[2] = 1000
	fmt.Println(a)
}
