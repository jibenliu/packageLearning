package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)
	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("First observable:%d\n", i)
	})
	time.Sleep(3 * time.Second)
	fmt.Println("before subscribe second observer")
	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("Second observer: %d\n", i)
	})

	time.Sleep(3 * time.Second)
}
