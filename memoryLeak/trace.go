package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

var mlock sync.RWMutex
var wg sync.WaitGroup

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go gets()
	}

	wg.Wait()
}

func gets() {
	for i := 0; i < 100000; i++ {
		get(i)
	}
	wg.Done()
}

func get(i int) {
	beginTime := time.Now()
	mlock.RLock()
	tmp1 := time.Since(beginTime).Nanoseconds() / 1000000
	if tmp1 > 100 { // 超过100ms就打印出来
		fmt.Println("fuck here")
	}
	mlock.RUnlock()
}


// go run main.go 2> trace.out
// go tool trace trace.out

// 			View trace：查看跟踪
//			Goroutine analysis：Goroutine 分析
//			Network blocking profile：网络阻塞概况
//			Synchronization blocking profile：同步阻塞概况
//			Syscall blocking profile：系统调用阻塞概况
//			Scheduler latency profile：调度延迟概况
//			User defined tasks：用户自定义任务
//			User defined regions：用户自定义区域
//			Minimum mutator utilization：最低 Mutator 利用率
