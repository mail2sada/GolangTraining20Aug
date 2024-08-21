package main

import "fmt"

type Addr struct {
	area    string
	pincode int
}

func (a Addr) Print1() {
	fmt.Println(a.area)
	fmt.Println(a.pincode)
}

type Emp1 struct {
	eid  int
	name string
	Addr
}

func (e Emp1) Print() {
	fmt.Println(e.eid)
	fmt.Println(e.name)
	e.Addr.Print1()
}

func main() {
	fmt.Println("Demo Methods")

	e := Emp1{eid: 1, name: "Anand", Addr: Addr{area: "HSR", pincode: 560102}}

	e.Print()
	e.Print1()
}
