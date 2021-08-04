package main

import (
	"context"
	"fmt"
	"time"
)

//gen 函数在单独的 Goroutine 中生成整数并将它们发送到返回的通道，gen 的调用者在使用生成的整数之后需要取消上下文，以免 gen 启动的内部 Goroutine 发生泄漏
func showContextTime() {
	start := time.Now() // 获取当前时间
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // return结束该goroutine，防止泄露
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		//fmt.Println(n)
		if n >= 100000 {
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Println("函数1执行完成耗时：", elapsed)
}

//显示普通循环耗时
func showLoopTime() {
	start := time.Now()
	for n := 0; n < 100000; n++ {
	}
	elapsed := time.Now().Sub(start)
	fmt.Println("函数2执行完成耗时：", elapsed)
}
func main() {
	showContextTime()
	showLoopTime()
}
