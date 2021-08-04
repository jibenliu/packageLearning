package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Range(0, 5)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
