package main

import "fmt"

type Address struct {
	Area    string
	Street  string
	City    string
	PinCode int32
}

type Employee struct {
	EmpID       int64
	Name        string
	Designation string
	Department  string
	Salary      float64
	Addr        Address
}

func main() {
	fmt.Println("Demo: structs")

	e1 := Employee{EmpID: 12345,
		Name:        "Anil kumar",
		Designation: "Software engineer",
		Department:  "RAN",
		Salary:      10000,
		Addr: Address{Area: "HSR",
			Street:  "27th",
			City:    "Bangalore",
			PinCode: 560102}}

	fmt.Println(e1)

}
