package main

import "fmt"

func main() {
	fmt.Println("Demo: String concat")

	s1 := "Hello"
	s2 := " World"

	s := s1 + s2

	fmt.Println(s)

	s += s1

	fmt.Println(s)

}
