/*
	Square

Generator	Square	Display

	Square
*/
package main

import (
	"fmt"
	"sync"
)

func GeneratorW(wg *sync.WaitGroup, elem ...int) chan int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for _, n := range elem {
			ch <- n
		}
	}()
	return ch
}

func SquareW(wg *sync.WaitGroup, input chan int) chan int {
	output := make(chan int)
	wg.Add(1)
	localWg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		localWg.Add(1)
		go func() {
			defer localWg.Done()
			for n := range input {
				square := n * n
				output <- square
			}
		}()
	}

	go func() {
		localWg.Wait()
		close(output)
		wg.Done()
	}()

	return output
}

func DisplayW(wg *sync.WaitGroup, ch chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range ch {
			fmt.Println(n)
		}
	}()
}

func main() {
	fmt.Println("Demo:Workerpool, fanout, fanin")

	wg := sync.WaitGroup{}

	DisplayW(&wg, SquareW(&wg, SquareW(&wg, GeneratorW(&wg, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))))
	wg.Wait()

}
