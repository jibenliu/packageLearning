package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
		next <- rxgo.Error(errors.New("unknown"))
		next <- rxgo.Of(4)
		next <- rxgo.Of(5)
	}})

	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error is:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}

	ob := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
		next <- rxgo.Error(errors.New("unknown"))
	},
		func(ctx context.Context, next chan<- rxgo.Item) {
			next <- rxgo.Of(4)
			next <- rxgo.Of(5)
		}})

	ch1 := ob.Observe()
	for item := range ch1 {
		if item.Error() {
			fmt.Println("error is:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}
