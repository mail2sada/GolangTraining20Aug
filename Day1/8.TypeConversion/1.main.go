package main

import "fmt"

func main() {
	s1, s2, s3 := 90, 98, 99

	avg := float64(s1+s2+s3) / 3

	rounded := int(avg)

	fmt.Println(avg, rounded)
}
