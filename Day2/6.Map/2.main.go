package main

import "fmt"

func main() {
	fmt.Println("Demo: Maps")

	mp := map[int]string{1: "One", 2: "Two", 3: "Three", 4: "Four"}

	for key, value := range mp {
		fmt.Println(key, value)
	}
}
