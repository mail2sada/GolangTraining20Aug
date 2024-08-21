package main

import "fmt"

func main() {
	fmt.Println("Demo: functions")

	PrintMessage("Hello")

	PrintMessage("hi")

	PrintMessage("Welcome")
}

func PrintMessage(msg string) {
	fmt.Println(msg)
}
