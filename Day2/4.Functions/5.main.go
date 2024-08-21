package main

import "fmt"

func CalculateStats(s1, s2, s3, s4 int) (int, float64) {

	tot := s1 + s2 + s3 + s4

	avg := float64(tot) / 4

	return tot, avg
}

func main() {
	fmt.Println("Demo: Function multiple returns")

	sum, avg := CalculateStats(100, 85, 99, 33)

	fmt.Println(sum, avg)

	_, avg = CalculateStats(1000, 85, 99, 33)

	fmt.Println(avg)

	sum, _ = CalculateStats(1000, 850, 99, 33)
	fmt.Println(sum)

}
