package main

import "fmt"

func main() {
	done := make(chan int) //无缓存的管道
	go func() {
		fmt.Printf("你好，世界\n")
		<-done
	}()
	done <- 1
}
