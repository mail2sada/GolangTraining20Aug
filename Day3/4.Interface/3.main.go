package main

import "fmt"

type Student struct {
	name string

	rollno int
}

func (s Student) String() string {
	fmt.Println("This from stringer")
	//return "this print is from out method" + s.name

	return fmt.Sprintln("11this print is from out method", s.name)
}

func main() {
	fmt.Println("Demo interface")

	s := Student{name: "Anand", rollno: 10}
	fmt.Println(s)

	str := fmt.Sprintln(s)
	fmt.Println(str)
}
