package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"sync/atomic"
	"time"
	"unsafe"
)

var query = "data"
var matches int

var workerCount int64 = 1
var maxWorkerCount = runtime.NumCPU()
var ch = make(chan struct{}, maxWorkerCount)
var searchRequest = make(chan string)
var workerDone = make(chan struct{})
var foundMatch = make(chan struct{})

func main() {
	start := time.Now()
	go search("/", true)
	waitForWorkers()
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}

func waitForWorkers() { //g0里面运行
	for {
		select {
		case path := <-searchRequest: //收到查询请求
			atomic.AddInt64(&workerCount, 1) //统计已用cpu数量
			ch <- struct{}{}                 //阻塞g0
			go search(path, true)            //开辟新g完成查询
		case <-workerDone: //g0已完成操作
			atomic.AddInt64(&workerCount, -1) //统计已用cpu数量
			fmt.Println("done chan workerCount:", workerCount)
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
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
				//fmt.Println("workerCount:", workerCount)
				if *(*int)(unsafe.Pointer(&workerCount))-1 < maxWorkerCount { //有空闲g即发送信号，让g0继续创建goroutine
					fmt.Println("now workerCount:", workerCount-1)
					searchRequest <- dir
				} else {
					search(dir, false) //所有g都忙碌，则在当前g继续查找
				}
			}
		}
		if master {
			workerDone <- struct{}{}
			<-ch
		}
	}
}
