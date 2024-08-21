package main

import "fmt"

func main() {
	fmt.Println("Map demo")

	mp := make(map[int]string)

	mp[1] = "One"
	mp[2] = "Two"
	mp[3] = "Three"
	mp[4] = "Four"

	for key, value := range mp {
		fmt.Println(key, value)
	}
	delete(mp, 2)
	for key, value := range mp {
		fmt.Println(key, value)
	}

}
