package main

import (
	"fmt"
	"reflect"
)

type Color int

const (
	Red   Color = 0
	Green Color = 1
	Blue  Color = 2
)

// func (c Color) String() string {
// 	switch c {
// 	case 0:
// 		return "red"
// 	case 1:
// 		return "green"
// 	case 2:
// 		return "blue"
// 	default:
// 		return ""
// 	}
// }

func main() {
	fmt.Println("Reflection with fmt package")

	r := Red
	b := Blue
	g := Green

	CheckStringerInterface(r)
	CheckStringerInterface(b)
	CheckStringerInterface(g)
	CheckStringerInterface(0)
	CheckStringerInterface("hello")
	CheckStringerInterface(3.142)

}

func CheckStringerInterface(v any) {
	fmt.Println("Implements Stringer:", reflect.TypeOf(v).Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()))
}
