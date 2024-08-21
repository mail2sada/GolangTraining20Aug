package main

import "fmt"

type Student struct {
	Roll_no int
	Name    string
}

func main() {
	fmt.Println("Demo: struct slice")

	stdslice := []Student{Student{Roll_no: 1, Name: "a"},
		Student{Roll_no: 2, Name: "B"},
		Student{Roll_no: 3, Name: "c"}}

	for i, v := range stdslice {
		fmt.Println(i)
		fmt.Println(v)
	}
}
