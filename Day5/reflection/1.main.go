package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Student struct {
	Roll int    `json:"roll_no"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Demo json encoding")
	str := `{
		"roll_no":1,
		"name":"Anand Kumar"
	}`

	std := Student{}

	err := json.Unmarshal([]byte(str), &std)
	fmt.Println(std, err)
	GetDetails(std)

}

func GetDetails(tp interface{}) {
	fmt.Println(reflect.TypeOf(tp))
	t := reflect.TypeOf(tp)
	v := reflect.ValueOf(tp)
	fmt.Println(v)

	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		name := t.Field(i).Name
		tp := t.Field(i).Type

		fmt.Println(tag, name, tp)
	}
}
