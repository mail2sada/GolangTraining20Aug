//Write a function circle to calculate area and circu

package main

import "fmt"

const PI = 3.14159

func Circle(radius float64) (float64, float64) {
	area := PI * radius * radius

	circum := 2 * PI * radius

	return area, circum
}

func main() {

	fmt.Println("Demo: Circles")

	a, c := Circle(2.5)

	fmt.Println(a, c)
}
