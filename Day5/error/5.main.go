package main

import "fmt"

type CustomError struct {
	StatusCode    int
	StatusMessage string
}

func (ce CustomError) Error() string {
	fmt.Println("Custome function Error() is used")
	return fmt.Sprintln("StatusCode:", ce.StatusCode, "StatusMessage:", ce.StatusMessage)
}

func GenerateError(key int) error {

	switch key {
	case 0:
		ce := CustomError{StatusCode: 400, StatusMessage: "Not found"}
		return ce
	case 1:
		ce := CustomError{StatusCode: 404, StatusMessage: "Bad Request"}
		return ce
	case 2:
		ce := CustomError{StatusCode: 503, StatusMessage: "Service Unavailable"}
		return ce
	default:
		return nil
	}
}

func main() {
	fmt.Println("Custom error!!!")

	e := GenerateError(0)
	if e != nil {
		fmt.Println(e)
		ce := e.(CustomError)
		fmt.Println(ce.StatusCode)
		fmt.Println(ce.StatusMessage)
	}
}
