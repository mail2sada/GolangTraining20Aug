package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Delete file")

	err := os.Remove("./test.tx")
	if err != nil {
		fmt.Println("Received error", err)
		return
	}
	fmt.Println("File deleted")
}
