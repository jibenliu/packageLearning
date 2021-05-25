package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}

	var ch = make(chan int)
	<-ch
}

/**
因为一开始就设置了只有一个 P，所以 for 循环里面“生产”出来的 goroutine 都会进入到 P 的 runnext 和本地队列，而不会涉及到全局队列
每次生产出来的 goroutine 都会第一时间塞到 runnext，而 i 从 1 开始，runnext 已经有 goroutine 在了，所以这时会把 old goroutine 移动 P 的本队队列中去，
再把 new goroutine 放到 runnext。之后会重复这个过程……
因此这后当一次 i 为 9 时，新 goroutine 被塞到 runnext，其余 goroutine 都在本地队列。
之后，main goroutine 执行了一个读 channel 的语句，这是一个好的调度时机：main goroutine 挂起，运行 P 的 runnext 和本地可运行队列里的 goroutine

而我们又知道，runnext 里的 goroutine 的执行优先级是最高的，因此会先打印出 9，接着再执行本地队列中的 goroutine 时，按照先进先出的顺序打印：0, 1, 2, 3, 4, 5, 6, 7, 8
 */