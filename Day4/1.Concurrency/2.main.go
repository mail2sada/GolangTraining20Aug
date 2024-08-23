package main

import (
	"fmt"
	"time"
)

func FirstGoroutine() {
	fmt.Println("I am running this in parallel")
}

func main() {
	fmt.Println("Demo concurrency")

	go FirstGoroutine()

	fmt.Println("Exiting main")

	time.Sleep(1 * time.Millisecond)
}
