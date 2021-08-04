package main

import (
	"fmt"
	"time"
)

func main() {
	foo()
	fmt.Println("ending")
}

func foo() {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("hello")
		default:
			return
		}
	}
}
