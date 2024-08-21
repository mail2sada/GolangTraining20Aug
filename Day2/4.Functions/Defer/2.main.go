package main

import "fmt"

func TestingDeferCode() {
	fmt.Println("We are in TestingDeferCode")      //3
	defer fmt.Println("This is defered statement") //5
	fmt.Println("TestingDeferCode exit")           //4
}

func main() {
	fmt.Println("Demo defer") //1
	defer TestingDeferCode()
	fmt.Println("Exiting main") //2
}
