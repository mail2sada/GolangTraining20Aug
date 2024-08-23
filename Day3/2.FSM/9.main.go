package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("WalkDir")

	var count int
	fsys := os.DirFS("/Users/sadanandd/Documents")
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			count++
		}
		fmt.Println(p)
		return nil
	})
	fmt.Println(count)
}
