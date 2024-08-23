package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Creating files")
	file, err := os.Create("./test.tx")
	if err != nil {
		fmt.Println("Could not create file 12334", err)
		return
	}
	defer file.Close()

	file.Write([]byte("This is our test file 123456"))
	fmt.Println("Wrote fle exiting")
}
