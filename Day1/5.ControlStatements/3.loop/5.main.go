package main

import "fmt"

func main() {
	fmt.Println("Demo: for loops")

	var a = 0

	for {
		a++
		if a%2 == 0 {
			fmt.Println(a)
		}
		if a == 99 {
			continue
		}
		if a >= 99 {
			break
		}
	}

}
