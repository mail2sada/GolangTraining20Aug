package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Demo: Unknown json format")
	var json_data = `{
		"roll_no": 1,
		"name": "Anil",
		"class":5,
		"address": "HSR Bangalore"
	}`
	mp := map[string]any{}

	err := json.Unmarshal([]byte(json_data), &mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mp)

}
