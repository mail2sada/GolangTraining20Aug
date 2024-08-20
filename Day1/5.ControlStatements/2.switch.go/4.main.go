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
	case Monday:
		fmt.Println("Beggining of the week")
		fallthrough
	case Tuesday:
		fmt.Println("Week is started")
	case Wednsday:
		fmt.Println("Mid of the week")
		fallthrough
	case Thursday:
		fmt.Println("Nearing weekend")
	case Friday:
		fmt.Println("Weekend is inch away")
		fallthrough
	case Saturday:
		fmt.Println("This is weekend")
		fallthrough
	case Sunday:
		fmt.Println("In weekend")
	default:
		fmt.Println("Don't know")
	}
}
