package main

import (
	"net/http"
	_ "net/http/pprof"
)

func funcA() []byte {
	a := make([]byte, 10*1024*1024)
	return a
}

func funcB() ([]byte, []byte) {
	a := make([]byte, 10*1024*1024)
	b := funcA()
	return a, b
}

func funcC() ([]byte, []byte, []byte) {
	a := make([]byte, 10*1024*1024)
	b, c := funcB()
	return a, b, c
}

func main() {
	for i := 0; i < 5; i++ {
		funcA()
		funcB()
		funcC()
	}
	http.ListenAndServe("0.0.0.0:9999", nil)
}

// runtime/pprof 采集程序（非 Server）的运行数据进行分析
// net/http/pprof：采集 HTTP Server 的运行时数据进行分析

// 在浏览器中输入http://ip:8899/debug/pprof/可以看到一个汇总页面，
	// cpu（CPU Profiling）: $HOST/debug/pprof/profile，默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件
	// block（Block Profiling）：$HOST/debug/pprof/block，查看导致阻塞同步的堆栈跟踪
	// goroutine：$HOST/debug/pprof/goroutine，查看当前所有运行的 goroutines 堆栈跟踪
	// heap（Memory Profiling）: $HOST/debug/pprof/heap，查看活动对象的内存分配情况
	// mutex（Mutex Profiling）：$HOST/debug/pprof/mutex，查看导致互斥锁的竞争持有者的堆栈跟踪
	// threadcreate：$HOST/debug/pprof/threadcreate，查看创建新OS线程的堆栈跟踪


// 动态获取 go tool pprof -alloc_space/-inuse_space http://ip:9999/debug/pprof/heap
 // -inuse_space：分析应用程序的常驻内存占用情况
 // -alloc_objects：分析应用程序的内存临时分配情况
	// flat：给定函数上运行耗时
	// flat%：同上的 CPU 运行耗时总比例
	// sum%：给定函数累积使用 CPU 总比例
	// cum：当前函数加上它之上的调用运行总耗时
	// cum%：同上的 CPU 运行耗时总比例

// 输出文件 go tool pprof -inuse_space -cum -svg http://ip:8899/debug/pprof/heap > heap_inuse.svg