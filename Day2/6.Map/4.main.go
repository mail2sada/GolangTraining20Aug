package main

import "fmt"

func main() {
	fmt.Println("Demo maps")

	mp := map[int]int{1: 1, 2: 10, 3: 100, 4: 1000, 5: 10000}

	fmt.Println(mp)

	for key, _ := range mp {
		delete(mp, key)
	}

	fmt.Println(mp)
}
