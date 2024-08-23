package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading file line by line")

	file, err := os.Open("./first.txt")

	if err != nil {
		fmt.Println("Received error", err)
		return
	}
	defer file.Close()
	data := make([]byte, 10, 10)
	size := 0
	for {
		n, err := file.Read(data)
		if err != nil {
			break
		}
		size += n
		fmt.Print(string(data))
	}
	fmt.Println()
	fmt.Println(size)

	n, e := file.ReadAt(data, 10)

	fmt.Println(n, e)
	fmt.Println(string(data))
}
