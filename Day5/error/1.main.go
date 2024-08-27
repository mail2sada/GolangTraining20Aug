package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("temp.txt")
	if err != nil {
		fmt.Println("We received an error", err)
	} else {

		fmt.Println("File open successful")
		file.Close()
	}
}
