package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Demo: strings package")

	s := "Hello world"

	sslice := strings.Split(s, " ")

	fmt.Println(sslice)

	for _, v := range sslice {
		fmt.Println(v)
	}

	s1 := "141"

	a, _ := strconv.Atoi(s1)

	fmt.Println(a)
}
