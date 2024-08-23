package main

import (
	"fmt"

	"github.com/mail2sada/training/math"
)

func main() {
	fmt.Println("Demo usage of published package")
	a := math.Add(20, 30)

	b := math.Sub(100, 50)

	fmt.Println(a, b)
}
