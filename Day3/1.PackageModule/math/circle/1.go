package circle

import "fmt"

func PrintValueOfUnexported() {
	fmt.Println(unExported)
	printUsage()
}
