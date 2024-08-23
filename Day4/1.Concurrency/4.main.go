package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Demo: Concurrency")

	var wg sync.WaitGroup = sync.WaitGroup{}

	wg.Add(1)

	data := 0

	go func() {
		defer wg.Done()

		for i := 0; i < 2000000; i++ {
			data++
		}
	}()
	wg.Wait()
	fmt.Println("Value of counter is ", data)
}
