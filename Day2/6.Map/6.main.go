package main

import "fmt"

func main() {
	fmt.Println("Demo: map of maps")

	var mp map[string]map[int]int = map[string]map[int]int{"A": map[int]int{1: 1, 2: 10},
		"B": map[int]int{3: 100, 4: 1000}}

	for key, value := range mp {
		fmt.Println(key)
		//fmt.Println(value)
		for k, v := range value {
			fmt.Println(k)
			fmt.Println(v)
		}
	}
}
