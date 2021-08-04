package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 1)
	ch <- 10
	go func(i <-chan int) {
		for x := range i {
			fmt.Printf("Process %d\n", x)
		}
	}(ch)

	time.Sleep(time.Second)
}
