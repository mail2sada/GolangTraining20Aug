package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
}

func (a Student) UpdateStudent(rollno int, name string) {
	a.RollNo = rollno
	a.Name = name
	fmt.Println(a)
}

func (a Student) PrintStudent() {
	fmt.Println("Student details...")
	fmt.Println("Roll no", a.RollNo)
	fmt.Println("Name", a.Name)
}

func main() {
	fmt.Println("Demo: Methods")

	std := Student{RollNo: 1, Name: "Anand"}

	std.PrintStudent()

	std.UpdateStudent(10, "Hello")

	fmt.Println(std)
}
