package main

import "fmt"

func main() {
	done := make(chan int, 10) //带缓存的10个管道

	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Printf("你好，世界%d\n", j)
			done <- 1
		}()
	}
	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}
