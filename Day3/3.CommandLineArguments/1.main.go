package main

import (
	"fmt"
	"os"
)

func main() {
	// ./1.main -name abc.txt -n 10
	fmt.Println("Demo command line arguments")
	fmt.Println(os.Args)
	args := os.Args[1:]
	fmt.Println(args)
}
