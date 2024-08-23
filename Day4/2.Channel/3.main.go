package main

import (
	"fmt"
	"time"
)

func ComputeSum(s, l int, output chan int) {
	defer close(output)
	sum := 0
	for i := s; i <= l; i++ {
		sum += i
		output <- sum
	}
}

func PrintChanStats(ch chan int) {
	for {
		fmt.Println("len:", len(ch))
		fmt.Println("cap:", cap(ch))
		time.Sleep(1 * time.Microsecond)
	}
}

func main() {
	fmt.Println("Demo: Buffered channels")
	ch := make(chan int, 5)

	go ComputeSum(10, 100, ch)

	go PrintChanStats(ch)
	for n := range ch {
		fmt.Println(n)
	}
}
