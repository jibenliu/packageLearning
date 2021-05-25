package main

import "fmt"

var (
	aVar int
	bVar func() bool
	cVar struct {
		d int
	}
)

func main() {
	aVar = 1
	fmt.Println(aVar)
	fmt.Println(bVar)
	fmt.Println(cVar)
}
