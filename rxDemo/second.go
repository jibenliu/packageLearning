package main

import (
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Just(1, 2, errors.New("unknown"), 3, 4, 5, 6.2)()
	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error is :", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}
