package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading files")

	data, err := os.ReadFile("/Users/sadanandd/Documents/GitHub/GolangTraining20Aug/Day3/2.FSM/1.main.go")
	if err != nil {
		fmt.Println("Received error", err)
		return
	}
	fmt.Println("Contents of file are", string(data))

}
