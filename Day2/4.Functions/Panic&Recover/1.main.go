package main

import "fmt"

func main() {
	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println("We paniced", rec)
		}
	}()
	fmt.Println("Demo: Panic")
	panic("We are panicing here")
	fmt.Println("Hello")
}
