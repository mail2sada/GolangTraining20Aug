package main

import "fmt"

func Concat(s1 string, s2 string) string {
	s := s1 + s2

	return s
}

func main() {
	fmt.Println("Demo: functions")

	a := Concat("Hello", "Hi")

	fmt.Println(a)

	fmt.Println(Concat(a, "World"))

	b := Concat(Concat("Abc", "Def"), Concat("Ghi", "Klm"))

	fmt.Println(b)

}
