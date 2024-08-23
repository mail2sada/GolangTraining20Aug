package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Exploring time package...")
	t := time.Now()
	fmt.Println("Current time is:", t)
	fmt.Println(t.Date())
	fmt.Println(t.Month())
	fmt.Println(t.Year())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Weekday())

}
