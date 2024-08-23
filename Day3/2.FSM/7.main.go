package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("Working with Directory")

	err := os.Mkdir("./root", fs.ModeDir)
	if err == nil {
		fmt.Println("Dir created")
	} else {
		fmt.Println("Failed to create dir", err)
	}

	err = os.MkdirAll("./1/2/3/4/5/6/7/8/9/10", os.ModeDir)
	if err == nil {
		fmt.Println("Dir created", err)
	} else {
		fmt.Println("Failed to create dir", err)
	}

	

}
