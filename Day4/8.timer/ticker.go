package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: tickers")

	t := time.NewTicker(500 * time.Millisecond)
	go func() {
		for n := range t.C {
			fmt.Println(n)
		}
	}()
	time.Sleep(10 * time.Second)
	t.Stop()
}
