package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var num = 0

func main() {
	n := runtime.NumCPU()
	fmt.Println(n)
	runtime.GOMAXPROCS(n - 1) //设置cpu运行的数目
	mu := &sync.Mutex{}
	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			num += 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("num:", num) // 如果不加锁这里的num的值会是一个随机数而不是10000
}
