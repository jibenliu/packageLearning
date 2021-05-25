package main

import (
	"fmt"
	"runtime"
	_ "runtime/pprof"
)

func main() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)

	intChan <- 1
	stringChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		panic(value)
	}
}
