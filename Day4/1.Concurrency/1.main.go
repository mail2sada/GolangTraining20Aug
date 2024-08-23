package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Concurrency")

	go fmt.Println("I want execute this in seperate go routine")
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Exiting main")

}
