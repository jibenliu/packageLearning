package main

import (
	"context"
	"fmt"
	"time"
)

func monitor3(ctx context.Context, number int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("监控器%v，监控结束。\n", number)
			return
		default:
			fmt.Printf("监控器%v，正在监控中...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx01, cancel := context.WithCancel(context.Background())
	ctx02, cancel := context.WithDeadline(ctx01, time.Now().Add(1*time.Second))

	defer cancel()

	for i := 1; i <= 5; i++ {
		go monitor3(ctx02, i)
	}

	time.Sleep(5 * time.Second)
	if ctx02.Err() != nil {
		fmt.Println("监控器取消的原因: ", ctx02.Err())
	}

	fmt.Println("主程序退出！！")
}
