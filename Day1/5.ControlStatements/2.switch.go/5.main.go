package main

import "fmt"

func main() {
	fmt.Println("Switch case")

	var a = 100

	switch {
	case a > 100:
		fmt.Println("greater than 100")
	case a < 100:
		fmt.Println("greater than 100")
	case a == 100:
		fmt.Println("equal to 100")
	}
}
