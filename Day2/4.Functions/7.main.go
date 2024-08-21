//Named returns

package main

import "fmt"

func Square(side float64) (area float64, circ float64) {

	area = side * side

	circ = 4 * side

	return
}

func main() {
	fmt.Println("Demo: Named return")

	a, c := Square(45.0)

	fmt.Println(a, c)
}
