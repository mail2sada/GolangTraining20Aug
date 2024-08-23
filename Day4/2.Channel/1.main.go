package main

import "fmt"

func main() {
	fmt.Println("Demo channels")

	var ch chan int = make(chan int)
	fmt.Println("Triggrting goroutine")
	go func(c chan int) {
		fmt.Println("Inside goroutine")
		var countr = 0
		for i := 0; i < 10000; i++ {
			countr++
		}
		fmt.Println("Writing to channel")
		c <- countr
	}(ch)

	fmt.Println("Reading value from channel")
	value := <-ch

	fmt.Println("Value of counter is", value)
}
