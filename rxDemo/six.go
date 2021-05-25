package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {
	observable := rxgo.Interval(rxgo.WithDuration(1 * time.Second))
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	// never running
	t := time.NewTicker(1 * time.Second)
	var count int
	for range t.C {
		fmt.Println("new ticker:", count)
		count++
	}
}
