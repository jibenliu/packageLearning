package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				fmt.Print(runtime.NumGoroutine())
				o <- true
				break
			}
		}
	}()
	<-o
}
