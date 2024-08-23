package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Timers again")
	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(3*time.Second, func() {
		defer wg.Done()
		fmt.Println("Timer has expired..")
	})
	wg.Wait()
}
