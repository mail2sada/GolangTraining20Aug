package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg = sync.WaitGroup{}

func IncrementCounter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		counter++
	}
}

func main() {
	fmt.Println("Demo: Goroutines")

	wg.Add(1)
	go IncrementCounter()
	wg.Add(1)

	go IncrementCounter()
	wg.Add(1)

	go IncrementCounter()
	wg.Add(1)

	go IncrementCounter()
	wg.Add(1)

	go IncrementCounter()
	wg.Wait()

	fmt.Println("Counter is incremented to", counter)
}
