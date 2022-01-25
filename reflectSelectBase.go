package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Generate the list of channel
	chs := make([]<-chan int, 0)
	for i := 0; i < 4; i++ {
		ch := make(chan int)
		go func(i int) { ch <- i }(i)
		chs = append(chs, ch)
	}

	// Convert the list of cases(SelectCase)
	cases := make([]reflect.SelectCase, len(chs))
	for i, ch := range chs {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
			// If Dir is SelectRecv, Send must be a zero Value
			// Send: ...
		}
	}

	// Wait until the channel receives the value (reflect.Select)
	for i := 0; i < 4; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if ok {
			fmt.Printf("chosen: %d, recv: %v\n", chosen, recv)
		}
	}
}
