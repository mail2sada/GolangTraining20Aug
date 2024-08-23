package spreadsheet

import "fmt"

type Sheet struct {
	sheetname string

	rows, column int
}

func (d Sheet) Print() {
	fmt.Println(d.sheetname, d.rows, d.column, "Sheet: Print")
}

func (d Sheet) PageSetup() {
	fmt.Println(d.sheetname, d.rows, d.column, "Sheet: PageSetup")

}

func (d Sheet) Preview() {
	fmt.Println(d.sheetname, d.rows, d.column, "Sheet: Preview")
}
