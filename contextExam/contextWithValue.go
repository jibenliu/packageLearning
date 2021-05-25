package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string // 定义一个key类型

	// f:一个从上下文中根据key取value的函数
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")

	// 创建一个携带key为k，value为"Go"的上下文
	ctx := context.WithValue(context.Background(), k, "Go")
	f(ctx, k)
	f(ctx, favContextKey("color"))

}
