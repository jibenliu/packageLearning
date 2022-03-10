package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

var query = "data"
var matches int64

var workerCount int64 = 0
var maxWorkerCount = runtime.NumCPU()
var ch = make(chan struct{}, maxWorkerCount) // 此处想实现的目标是同时存在NumCPU个goroutine来完成search目标（不能多于这么多个）

// 添加了default分支，所以chan可以设置成无缓冲的
var searchRequest = make(chan string, maxWorkerCount)
var workerDone = make(chan struct{}, maxWorkerCount)
var foundMatch = make(chan struct{}, maxWorkerCount)

func main() {
	start := time.Now()
	go search("/", true)
	waitForWorkers()
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}

func waitForWorkers() { //只在g0里面运行
	for {
		select {
		case path := <-searchRequest: //收到查询请求
			ch <- struct{}{} //因为带了缓冲，所以g0没有阻塞,最多有NumCPU个g存在，但是当数量超过该值的时候就阻塞在这里
			//fmt.Println("get create goroutine chan signal :", workerCount, "now goroutine number is :", runtime.NumGoroutine())
			workerCount++
			go search(path, true) //开辟新g完成查询
		case <-workerDone: //master g已完成操作
			workerCount--
			//fmt.Println("get done chan signal workerCount:", workerCount, "now goroutine number is :", runtime.NumGoroutine())
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			//fmt.Println("get matches goroutine chan signal :", matches, "now goroutine number is :", runtime.NumGoroutine())
			matches++
		}
	}
}

func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				foundMatch <- struct{}{}
			}
			if file.IsDir() {
				dir := path + "/" + name
				if path == "/" {
					dir = path + name
				}
				// 当前逻辑在非g0里面执行，所以用原子读
				if workerCount < int64(maxWorkerCount)-2 { //有空闲g即发送信号，让g0继续创建goroutine
					searchRequest <- dir
				} else {
					search(dir, false) //所有g都忙碌，则在当前g继续查找
				}
			}
		}
	}
	if master {
		workerDone <- struct{}{}
		<-ch
	}
}
