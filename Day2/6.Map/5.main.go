package main

import "fmt"

func main() {
	fmt.Println("Maps")

	mp := map[int]int{1: 1, 2: 10, 3: 100, 4: 1000, 5: 10000}

	fmt.Println(mp)
	//mp = nil
	mp = make(map[int]int)
	mp[1] = 1
	fmt.Println(mp)
}
