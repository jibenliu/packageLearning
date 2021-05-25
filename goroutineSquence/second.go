package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Hour)
}

/**
go 1.13 的 time 包会生产一个名字叫 timerproc 的 goroutine 出来，
它专门用于唤醒挂在 timer 上的时间未到期的 goroutine；
因此这个 goroutine 会把 runnext 上的 goroutine 挤出去。因此输出顺序就是：0, 1, 2, 3, 4, 5, 6, 7, 8, 9。

而 go 1.14 把这个唤醒的 goroutine 干掉了，取而代之的是，在调度循环的各个地方、sysmon 里都是唤醒 timer 的代码，timer 的唤醒更及时了，
但代码也更难看懂了。所以，输出顺序和第一个例子是一致的
 */