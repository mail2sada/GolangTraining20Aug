package main

import "fmt"

type Block struct {
	Try     func()
	Catch   func(any)
	Finally func()
}

func (b Block) Do() {
	defer b.Finally()
	defer func() {
		rec := recover()
		if rec != nil {
			b.Catch(rec)
		}
	}()
	b.Try()
}

func main() {
	fmt.Println("Demo try Catch")

	Block{
		Try: func() {
			fmt.Println("I am in try block")
			a := 100
			b := 0
			c := a / b
			fmt.Println(c)

		},
		Catch: func(e any) {
			fmt.Println("Got exception", e)
		},
		Finally: func() {
			fmt.Println("We are in finally block")
		},
	}.Do()

	fmt.Println("Exiting main")

}
