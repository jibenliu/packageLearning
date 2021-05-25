package main

import (
	"fmt"
	"sync"
	"time"
)

//互斥锁
func main() {
	var a = 0
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}

	//等待1s结束主程序
	//确保所有协程执行完
	time.Sleep(time.Second)
}
