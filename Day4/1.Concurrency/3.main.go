package main

import (
	"fmt"
	"time"
)

var counter = 0

func increment() {
	fmt.Println("Entering goroutine")
	for i := 0; i < 20000; i++ {
		counter++
	}
	fmt.Println("Exiting goroutine")
}

func increment1() {
	fmt.Println("Entering goroutine")
	for i := 0; i < 30000; i++ {
		counter++
	}
	fmt.Println("Exiting goroutine")
}

func main() {
	fmt.Println("Demo: Concurrency Goroutines")
	go increment()
	go increment1()
	go increment()
	time.Sleep(1 * time.Second)

	fmt.Println("Value of counter is", counter)
	fmt.Println("Exiting main...")

}
