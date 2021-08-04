package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {
	ch := make(chan rxgo.Item, 1)

	go func() {
		i := 0
		for range time.Tick(time.Second) {
			ch <- rxgo.Of(i)
			i++
		}
	}()

	observable := rxgo.FromChannel(ch).BufferWithTimeOrCount(rxgo.WithDuration(3*time.Second), 2)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
