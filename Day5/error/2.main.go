package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (res int, err error) {
	defer func() {
		rec := recover()

		if rec != nil {
			fmt.Println("Recovered panic", rec)
		}
	}()
	if b == 0 {
		err = errors.New("denominator can not be zero")
		err = fmt.Errorf("denominator can not be zero")
	} else {
		res = a / b
	}
	return
}

func main() {
	fmt.Println("Demo errors...")

	ans, e := Divide(10, 10)
	if e != nil {
		fmt.Println("We have received an error", e)
	} else {
		fmt.Println("Result is :", ans)
	}

	ans, e = Divide(10, 0)
	if e != nil {
		fmt.Println("We have received an error", e)
	} else {
		fmt.Println("Result is :", ans)
	}
}
