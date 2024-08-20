package main

import "fmt"

func main() {

	const (
		Sunday   = "Sunday"
		Monday   = "Monday"
		Tuesday  = "Tuesday"
		Wednsday = "Wednsday"
		Thursday = "Thursday"
		Friday   = "Friday"
		Saturday = "Saturday"
	)
	fmt.Println("Demo: switch case")

	var a = Friday
	switch a {
	case Monday, Tuesday:
		fmt.Println("Beggining of the week")

	case Wednsday, Thursday:
		fmt.Println("Mid of the week")

	case Friday, Saturday, Sunday:
		fmt.Println("Weekend is inch away")
	default:
		fmt.Println("Don't know")
	}
}
