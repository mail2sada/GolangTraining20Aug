package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time.Format & time.Parse")

	tm := time.Now()

	formatstr := "2006/Jan/02 **** **** 15:04:05"
	str := tm.Format(formatstr)
	fmt.Println(str)
}
