package main

import "fmt"

func main() {
	fmt.Println("Annonymous functions")

	func() {
		fmt.Println("Hello world!!!")
	}()

	
}
