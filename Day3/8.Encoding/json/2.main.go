package main

import (
	"encoding/json"
	"fmt"
)

type Student1 struct {
	RollNo  int    `json:"roll_no"`
	Name    string `json:"name"`
	Class   int    `json:"class"`
	Address string `json:"address"`
}

func main() {
	fmt.Println("Json marshaling")
	s := Student1{RollNo: 1, Name: "Anilkumar", Class: 10, Address: "HSR Bangalore"}

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	data, err = json.MarshalIndent(s, "\n", "\t")
	fmt.Println(string(data))

}
