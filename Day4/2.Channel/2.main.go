/*

 */

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func ComputeTotal(start, limit int, output chan int) {
	defer wg.Done()
	//defer close(output)
	sum := 0
	for i := start; i <= limit; i++ {
		sum += i
		output <- sum
	}
}

func main() {
	fmt.Println("Something Intresting with channels")

	ch := make(chan int)
	wg.Add(2)
	go ComputeTotal(100, 1000, ch)

	go ComputeTotal(5000, 100, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()
	for sum := range ch {
		fmt.Println(sum)
	}

}
