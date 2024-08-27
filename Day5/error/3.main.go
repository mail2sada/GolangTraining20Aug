package main

import (
	"errors"
	"fmt"
	"os"
)

func OpenFile(fname string) (f *os.File, err error) {
	f, err = os.Open(fname)
	if err != nil {
		fmt.Println("error:", err)
		err = fmt.Errorf("We have received an error %w", err)
	}
	return
}

func main() {
	fmt.Println("Some insights in error")

	file, e := OpenFile("./3.main.go")
	if e != nil {
		fmt.Println("Received an error", e)
	} else {
		fmt.Println("File opening is successful")
		file.Close()
	}

	file, e = OpenFile("./10.main.go")
	if e != nil {
		fmt.Println("Received an error", e)
		err := errors.Unwrap(e)
		fmt.Println(err)
	} else {
		fmt.Println("File opening is successful")
		file.Close()
	}
}
