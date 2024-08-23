package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Unix time or ephoc time")

	tm := time.Now()
	fmt.Println(tm)
	unixTime := tm.Unix()
	fmt.Println(unixTime)

	t := time.Unix(unixTime, 0)
	fmt.Println(t)
}
