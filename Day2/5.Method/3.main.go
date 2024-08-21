package main

import "fmt"

type Addr struct {
	area    string
	pincode int
}

func (a Addr) PrintAddress() {
	fmt.Println(a.area)
	fmt.Println(a.pincode)
}

type Emp struct {
	eid  int
	name string
	addr Addr
}

func (e Emp) PrintEmp() {
	fmt.Println(e.eid)
	fmt.Println(e.name)
	e.addr.PrintAddress()
}

func main() {
	fmt.Println("Demo Methods")

	e := Emp{eid: 1, name: "Anand", addr: Addr{area: "HSR", pincode: 560102}}

	e.PrintEmp()
}
