package main

import (
	"first/math"
	"first/math/circle"
	"first/math/square"
	"first/service"
	"fmt"
)

func CheckCirclePackage() {
	a, c := circle.CircleProperties(10.5)
	fmt.Println(a, c)

	fmt.Println(circle.PI)

	fmt.Println(circle.Exported)

	//fmt.Println(circle.unExported)
	fmt.Println("Printing value")
	circle.PrintValueOfUnexported()

}

func CheckMathPackage() {
	fmt.Println("Inside CheckMathPackage")
	s := math.AddNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(s)

	m := math.Multiply(10, 20)

	fmt.Println(m)

	d := math.Divide(100, 200)

	fmt.Println(d)

	diff := math.Subtract(500, 200)

	fmt.Println(diff)

	d = math.Divide(100, 0)
	fmt.Println(d)

}

func main() {
	fmt.Println("Demonstraiting package and modules")

	var c Config = Config{ServiceName: "ConfiService", Description: "Service handles and manages configuration"}

	serviceName := c.GetServiceName()
	description := c.GetDescription()

	fmt.Println("ServiceName is ", serviceName)
	fmt.Println("Description:", description)

	var s service.Service = service.Service{}
	// s.SetServiceName("Config service")
	// s.SetServiceID(1)
	// s.SetServiceDescription("Manages configurations")

	s = s.SetServiceName("Config Service").
		SetServiceID(1).
		SetServiceDescription("Manages configurations")

	name, id, desc := s.GetServiceDetails()

	fmt.Println(name, id, desc)

	var v service.Conf = service.Conf{ConfID: 1, ConfName: "Test"}

	fmt.Println("Conf struc from service")
	fmt.Println(v.ConfID, v.ConfName)

	CheckMathPackage()
	CheckCirclePackage()
	CheckSqurePackage()
}

func CheckSqurePackage() {
	a, c := square.SquareProperties(10)
	fmt.Println(a, c)
}
