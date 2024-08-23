/*
Generator numbers
Square
Display
*/

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func Generator(element ...int) chan int {
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		defer close(ch)
		for _, n := range element {
			ch <- n
			//write to channel
		}
	}()
	return ch
}

func Square(ch chan int) chan int {
	wg.Add(1)
	squarech := make(chan int)
	go func() {
		defer wg.Done()
		defer close(squarech)
		for n := range ch {
			squre := n * n
			squarech <- squre
			//Write this to another channel
		}
	}()
	return squarech
}

func Display(ch chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range ch {
			fmt.Println(n)
		}
	}()
}

func main() {
	fmt.Println("Concurrency Pattern:Pipeline")

	// gench := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// sqch := Square(gench)

	// Display(sqch)

	// Display(Square(Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)))
	// Display(Square(Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)))
	// Display(Square(Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)))

	Display(Square(Square(Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14))))

	wg.Wait()
}
