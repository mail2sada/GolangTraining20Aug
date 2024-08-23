package main

import (
	"fmt"
	"os"
)

func main() {
	f, e := os.Create("./first.txt")
	if e != nil {
		fmt.Println("received error", e)
		return
	}
	defer f.Close()

	fmt.Fprintln(f, "This is out first.txt file")
	fmt.Fprintln(f, " Line 2This is out first.txt file")

}
