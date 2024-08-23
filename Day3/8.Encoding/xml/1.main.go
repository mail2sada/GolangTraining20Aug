package main

import (
	"encoding/xml"
	"fmt"
)

type Student struct {
	RollNo int    `xml:"rollno"`
	Name   string `xml:"name"`
	Class  string `xml:"class"`
}

func main() {

	s := Student{RollNo: 1, Name: "Anil", Class: "10B"}

	data, err := xml.Marshal(s)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	data, err = xml.MarshalIndent(s, "\n", "\t")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}
