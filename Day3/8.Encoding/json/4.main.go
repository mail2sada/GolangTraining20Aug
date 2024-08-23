package main

import (
	"encoding/json"
	"fmt"
)

/*
{
	"roll_no": 1,
	"name": "Anil",
	"class":5,
	"address": "HSR Bangalore"
}
*/

type Student struct {
	RollNo  int    `json:"roll_no"`
	Name    string `json:"name"`
	Class   int    `json:"class"`
	Address string `json:"address"`
}

func (s Student) Print() {
	fmt.Println("RollNo", s.RollNo)
	fmt.Println("Name", s.Name)
	fmt.Println("Address", s.Address)
	fmt.Println("Class", s.Class)
}

var json_data = `[{
	"roll_no": 1,
	"name": "Anil",
	"class":5,
	"address": "HSR Bangalore"
},
{
	"roll_no": 2,
	"name": "Anil",
	"class":5,
	"address": "HSR Bangalore"
},
{
	"roll_no": 3,
	"name": "Anil",
	"class":5,
	"address": "HSR Bangalore"
}]`

func main() {
	fmt.Println("Demo json encoding")
	stdSlice := []Student{}
	//stud := Student{}
	err := json.Unmarshal([]byte(json_data), &stdSlice)
	if err != nil {
		fmt.Println("Received err", err)
		return
	}
	fmt.Println(stdSlice)

	for _, std := range stdSlice {
		std.Print()
	}

}
