package main

import "fmt"

func main() {
	done := make(chan int, 1)
	go func() {
		fmt.Printf("你好，世界\n")
		done <- 1
	}()
	<-done
}
