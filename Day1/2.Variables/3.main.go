package main

import "fmt"

func main() {
	fmt.Println("Demo:Variable declaration!!")
	var v1, v2, v3 int = 100, 200, 300

	fmt.Println(v1, v2, v3)

	var v4, v5, v6 = 100, 500, 800
	fmt.Println(v4, v5, v6)

	var v uint8 = 255

	fmt.Println(v)
}
