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

	var a = Sunday
	fmt.Printf("Data type of a %T\n", a)
	switch a {
	case Monday:
		fmt.Println("Beggining of the week")
	case Tuesday:
		fmt.Println("Week is started")
	case Wednsday:
		fmt.Println("Mid of the week")
	case Thursday:
		fmt.Println("Nearing weekend")
	case Friday:
		fmt.Println("Weekend is inch away")
	case Saturday:
		fmt.Println("This is weekend")
	case Sunday:
		fmt.Println("In weekend")
	}
}
