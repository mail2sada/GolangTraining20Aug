package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Timers")

	timer := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(10 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()
	fmt.Println(time.Now())
	tm := <-timer.C
	fmt.Println(tm)
	//fmt.Println(time.Now())

}
