package main

import "fmt"

var (
	ErrorTimeout            = fmt.Errorf("timeout")
	ErrorServiceUnavailable = fmt.Errorf("service unavailable")
	ErrorNotFound           = fmt.Errorf("not found")
)

func PerformOperation(input int) error {

	switch input {
	case 0:
		return ErrorTimeout
	case 1:
		return ErrorServiceUnavailable
	case 3:
		return ErrorNotFound
	}

	return nil
}

func main() {
	fmt.Println("Demo: defining errors")

	e := PerformOperation(3)
	switch e {
	case ErrorTimeout:
		fmt.Println("We have to retry")
	case ErrorServiceUnavailable:
		fmt.Println("Raise an alarm, update kpi")
	case ErrorNotFound:
		fmt.Println("Something wrong with input")
	default:
		fmt.Println("Operation executed...")
	}

	if e == nil {
		fmt.Println("Success")
	} else if e == ErrorTimeout {

	} else if e == ErrorServiceUnavailable {

	} else if e == ErrorNotFound {

	} else {
		//new error
	}
}
