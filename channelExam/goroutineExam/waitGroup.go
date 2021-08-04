package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			fmt.Println("hh")
			time.Sleep(time.Second)
			waitGroup.Done() // 减少需要等待goroutine的数量 相当于Add(-1)
		}()
	}

	waitGroup.Wait() // 执行阻塞，直到所有的需要等待的goroutine数量变成0
	fmt.Println("over")
}
