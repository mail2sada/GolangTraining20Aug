package main

import "fmt"

func main() {
	fmt.Println("Demo: Maps")

	var mp map[string]string = map[string]string{"A": "Apple", "B": "Ball", "C": "Cat"}

	fmt.Println(mp)

	mp["D"] = "Dog"
	fmt.Println(mp)

	v := mp["B"]
	fmt.Println(v)

	v = mp["E"]
	fmt.Println(v)
	v, ok := mp["E"]
	if !ok {
		fmt.Println("Not found")
	}
}
