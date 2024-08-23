package printer

type Printer interface {
	Print()
	PageSetup()
	Preview()
}

func Print(p Printer) {
	p.Print()
}

func PageSetup(p Printer) {
	p.PageSetup()
}

func Preview(p Printer) {
	p.Preview()
}
