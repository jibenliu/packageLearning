package main

import (
	"fmt"
	"time"
)

func monitor(ch chan bool, number int) {
	for {
		select {
		case v := <-ch:
			// 仅当 ch 通道被 close，或者有数据发过来(无论是true还是false)才会走到这个分支
			fmt.Printf("监控器%v，接收到通道值为：%v，监控结束。\n", number, v)
			return
		default:
			fmt.Printf("监控器%v，正在监控中...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	stopSignal := make(chan bool)
	for i := 1; i <= 5; i++ {
		go monitor(stopSignal, i)
	}

	time.Sleep(1 * time.Second)
	// 关闭所有 goroutine
	close(stopSignal)

	// 等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)

	fmt.Println("主程序退出！！")
}
