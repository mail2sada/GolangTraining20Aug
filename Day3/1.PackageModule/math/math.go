package math

import "fmt"

func AddNumbers(elements ...int) int {
	sum := 0

	for _, n := range elements {
		sum += n
	}
	return sum
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	defer func() {
		rec := recover()

		if rec != nil {
			fmt.Println("Received panic", rec)
		}
	}()
	return a / b
}

func Subtract(a, b int) (res int) {
	res = a - b
	return
}
