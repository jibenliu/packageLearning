package main

import (
	"fmt"
	"time"
)

//简单装饰器模式
func myFunc() {
	fmt.Println("hello world")
	time.Sleep(1 * time.Second)
}

func main() {
	fmt.Printf("Type: %T\n", myFunc)
}
