package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func readMemStats() {

	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func test() {
	//slice 会动态扩容，用slice来做堆内存申请
	container := make([]int, 8)

	log.Println(" ===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			readMemStats()
		}
	}

	log.Println(" ===> loop end.")
}

func main() {

	//启动pprof
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()

	log.Println(" ===> [Start].")

	readMemStats()
	test()
	readMemStats()

	log.Println(" ===> [force gc].")
	runtime.GC() //强制调用gc回收

	log.Println(" ===> [Done].")
	readMemStats()

	go func() {
		for {
			readMemStats()
			time.Sleep(10 * time.Second)
		}
	}()

	time.Sleep(3600 * time.Second) //睡眠，保持程序不退出
}


// http://127.0.0.1:10000/debug/pprof/heap?debug=1