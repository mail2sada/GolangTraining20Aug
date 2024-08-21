package main

import "fmt"

func main() {
	fmt.Println("Demo: defer functions")               //1
	defer fmt.Println("Call this later")               //3
	defer fmt.Println("Call this function also later") //3
	fmt.Println("We are exiting main")                 //2
}
