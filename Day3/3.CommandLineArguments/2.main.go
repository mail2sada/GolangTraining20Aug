package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Tool to read from a file with -name and -n lines")
	name := flag.String("name", "", "please enter the file name with option -name")
	line := flag.Int("n", 0, "Please enter how many lines you want print with option -n")
	flag.Parse()
	fmt.Println(*name, *line)

	if *name == "" || *line == 0 {
		flag.Usage()
	}

	data, err := os.ReadFile(*name)
	if err != nil {
		fmt.Println(err)
		return
	}
	strContents := string(data)
	lines := strings.Split(strContents, "\n")

	for n, v := range lines {
		if n == *line {
			break
		}
		fmt.Println(v)
	}
}
