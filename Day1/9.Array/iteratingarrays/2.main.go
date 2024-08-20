package main

import "fmt"

func main() {
	fmt.Println("Demo: Array iteration")

	var array = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
