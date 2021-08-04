package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 传递带有超时的上下文
	// 告诉阻塞函数在超时结束后应该放弃其工作。
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 终端输出"context deadline exceeded"
	}
}
