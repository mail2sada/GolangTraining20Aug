package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo ticker")

	a := func() {
		time.AfterFunc(500*time.Millisecond, a)
	}
}
