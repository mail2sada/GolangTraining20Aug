package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Time package")

	t := time.Now()
	fmt.Println(t)
	t = t.Add(45 * time.Minute)
	var d time.Duration = 10000 * time.Second
	fmt.Println("Duration is", d)
	fmt.Println(t)
}
