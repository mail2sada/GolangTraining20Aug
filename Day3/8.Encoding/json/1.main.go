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

var json_data = `{
	"roll_no": 1,
	"name": "Anil",
	"class":5,
	"address": "HSR Bangalore"
}`

func main() {
	fmt.Println("Demo json encoding")
	stud := Student{}
	err := json.Unmarshal([]byte(json_data), &stud)
	if err != nil {
		fmt.Println("Received err", err)
		return
	}
	fmt.Println(stud)

}
