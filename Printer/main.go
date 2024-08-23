package main

import (
	"fmt"
	"hpprint/doc"
	"hpprint/printer"
	"hpprint/spreadsheet"
)

func main() {
	fmt.Println("Demo: Interface with printer example")

	var d doc.Doc = doc.Doc{}

	var s spreadsheet.Sheet = spreadsheet.Sheet{}

	//var p printer.Printer

	printer.PageSetup(d)
	printer.Preview(d)
	printer.Print(d)

	printer.PageSetup(s)
	printer.Preview(s)
	printer.Print(s)

}
