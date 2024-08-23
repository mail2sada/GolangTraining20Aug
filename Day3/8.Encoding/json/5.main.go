package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	Pincode int    `json:"pincode"`
}

type Employee struct {
	Eid  int     `json:"eid"`
	Name string  `json:"name"`
	Addr Address `json:"address"`
}

func main() {
	fmt.Println("Nested json")

	e := Employee{Eid: 1, Name: "Anandkumar", Addr: Address{City: "Bangalore", Pincode: 560102}}

	data, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	data, err = json.MarshalIndent(e, "\n", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
