package main

import "fmt"

func main() {
	fmt.Println("Demo: nested if")

	var a, b = 300, 205

	if b > a {
		if b%10 == 0 {
			fmt.Println("Step at line 12")
		} else {
			fmt.Println("Step at line 14")
		}
	} else {
		fmt.Println("Step at line 17")
	}
}
