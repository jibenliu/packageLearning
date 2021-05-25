package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {
	observable := rxgo.Just(1, 2, 3)().Repeat(
		3, rxgo.WithDuration(1*time.Second),
	)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
