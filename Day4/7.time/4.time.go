package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo time")

	tm := time.Date(2024, 8, 23, 14, 0, 0, 0, &time.Location{})
	fmt.Println(tm)
}
