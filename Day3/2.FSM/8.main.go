package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Demo: Dir list")

	d, err := os.ReadDir("/Users/sadanandd/Documents/GitHub")
	if err != nil {
		fmt.Println("received error", err)
		return
	}

	//fmt.Println(d)

	for _, fnode := range d {
		fmt.Println(fnode.Name())
		finfo, _ := fnode.Info()
		fmt.Println(finfo.ModTime(), finfo.Mode(), finfo.Size())

	}
}
