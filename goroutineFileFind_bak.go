package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

var query = "data"
var matches int

var workerCount = 0
var maxWorkerCount = runtime.NumCPU()
var ch = make(chan struct{}, runtime.NumCPU())
var searchRequest = make(chan string)
var workerDone = make(chan struct{})
var foundMatch = make(chan struct{})

func main() {
	start := time.Now()
	workerCount = 1
	go search("/", true)
	waitForWorkers()
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}

func waitForWorkers() { //只在main中goroutine里面运行,故matches无并发修改
	for {
		select {
		case path := <-searchRequest: //收到查询请求
			fmt.Println("receive search request ", path)
			workerCount++
			ch <- struct{}{}
			go search(path, true) //开辟新g执行查询
		case <-workerDone: //master g已完成查询
			workerCount--
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
			FoundLoop:
				for {
					select {
					case foundMatch <- struct{}{}:
						break FoundLoop
					default:
					}
				}
			}
			if file.IsDir() {
				var dirPath = path + name + "/"
				if workerCount < maxWorkerCount {
				SearchLoop:
					for {
						select {
						case searchRequest <- dirPath:
							break SearchLoop
						default:
						}
					}
				} else {
					search(dirPath, false) //所有g都忙碌，则在当前g继续查找
				}
			}
		}
	}
	if master {
	DoneLoop:
		for {
			select {
			case workerDone <- struct{}{}:
				break DoneLoop
			default:
			}
		}
		<-ch
	}
}
