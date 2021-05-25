package main

import (
	"fmt"
	"time"
)

func newFunc() {
	fmt.Println("Hello world")
	time.Sleep(1 * time.Second)
}

func coolFunc(a func()) {
	a()
}

func main() {
	fmt.Printf("Type:%T\n", newFunc)
	coolFunc(newFunc)
}
