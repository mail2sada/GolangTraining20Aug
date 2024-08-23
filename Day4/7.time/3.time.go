package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time.Since")

	t1 := time.Now()
	time.Sleep(4 * time.Second)
	d := time.Since(t1)

	fmt.Println(d)
}
