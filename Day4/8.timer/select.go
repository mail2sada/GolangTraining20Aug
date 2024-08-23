package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Tickers")

	t1 := time.NewTicker(10 * time.Millisecond)
	t2 := time.NewTicker(100 * time.Millisecond)
	t3 := time.NewTicker(1000 * time.Millisecond)

	for {
		select {
		case a := <-t1.C:
			fmt.Println("Ticker1 ", a)
		case a := <-t2.C:
			fmt.Println("Ticker2 ", a)
		case a := <-t3.C:
			fmt.Println("Ticker3 ", a)

		}
	}
}
