package circle

import "fmt"

const PI = 3.14159

var Exported = 100

var unExported = 500

func CircleProperties(radius float64) (area, circ float64) {
	area = PI * radius * radius
	circ = 2 * PI * radius
	printUsage()
	return
}

func printUsage() {
	fmt.Println("We are priting usage string here")
}
