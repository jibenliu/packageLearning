package main

import (
	"fmt"
	"runtime"
	"sync"
)

//在 10 个 goroutine 中共享了变量 counter。每个 goroutine 执行完成后，会将 counter 的值加 1。因为 10 个 goroutine 是并发执行的，所以我们还引入了锁，也就是代码中的 lock 变量。每次对 n 的操作，都要先将锁锁住，操作完成后，再将锁打开
var counter int

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
