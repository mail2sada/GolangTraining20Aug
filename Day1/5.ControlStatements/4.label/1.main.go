package main

import "fmt"

func main() {
	fmt.Println("Demo: Labels")

	var a, b int
outer:
	a = 0
	b = 0
	for a <= 30 {
		a++
		b = 0
	inner:
		for b <= 20 {
			b++
			fmt.Println(b, a)
			if a == 5 {
				fmt.Println("continue outer")
				goto outer
			}
			// if b == 3 {
			// 	fmt.Println("continue inner")

			// 	continue inner
			// }

			if a == 10 {
				fmt.Println("break inner")

				break inner
			}

			if a == 20 {
				fmt.Println("break outer")

				break
			}
		}
	}
}
