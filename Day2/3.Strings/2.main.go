package main

import "fmt"

func main() {
	fmt.Println("Demo: string \nimmutablity")

	s := "Mavenir \ngolang training"

	for _, v := range s {
		fmt.Println(v)
		fmt.Printf("%c\n", v)
	}
	fmt.Println(s)

}
