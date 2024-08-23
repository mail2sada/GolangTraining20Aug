package doc

import "fmt"

type Doc struct {
	docName string

	contents string
}

func (d Doc) Print() {
	fmt.Println(d.docName, d.contents, "Docuemts: Print")
}

func (d Doc) PageSetup() {
	fmt.Println(d.docName, d.contents, "Docuemts: PageSetup")

}

func (d Doc) Preview() {
	fmt.Println(d.docName, d.contents, "Docuemts: Preview")
}
