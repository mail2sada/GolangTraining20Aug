package main

import "fmt"

type Employee1 struct {
	EmpID       int64
	Name        string
	Designation string
	Department  string
	Salary      float64
}

func main() {
	fmt.Println("Demo: structs")

	e1 := Employee1{EmpID: 12345,
		Name:        "Anil kumar",
		Designation: "Software engineer",
		Department:  "RAN",
		Salary:      10000}

	fmt.Println(e1)

}
