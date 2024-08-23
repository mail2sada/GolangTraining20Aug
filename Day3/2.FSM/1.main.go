package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("FSM ")

	file, err := os.Open("/Users/sadanandd/Documents/GitHub/GolangTraining20Aug/Day3/2.FSM/1.main.go")

	if err != nil {
		fmt.Println("We have got an error")
		return
	}
	defer file.Close()
	data := make([]byte, 1024, 2048)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println("Received error")
		return
	}

	fmt.Println("Contents of file", string(data), n, "bytes read")

}
