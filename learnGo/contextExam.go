package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return //return结束该goroutine，防止泄露
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //当我们取完需要的整数后调用cancel
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	d := time.Now().Add(50 * time.Microsecond)
	ctx, cancel = context.WithDeadline(context.Background(), d)
	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	//传递带有超时的上下文
	//高速阻塞函数在超时结束后应该放弃其工作
	ctx, cancel = context.WithTimeout(context.Background(), 50&time.Millisecond)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	type favContextKey string //定义一个key类型
	//f:一个从上下文中根据key取value的函数
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	k := favContextKey("language")
	//创建一个携带key为k，value为"Go"的上下文
	ctx = context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

	/**
	使用 Context 的注意事项：
	不要把 Context 放在结构体中，要以参数的方式显示传递；
	以 Context 作为参数的函数方法，应该把 Context 作为第一个参数；
	给一个函数方法传递 Context 的时候，不要传递 nil，如果不知道传递什么，就使用 context.TODO；
	Context 的 Value 相关方法应该传递请求域的必要数据，不应该用于传递可选参数；
	Context 是线程安全的，可以放心的在多个 Goroutine 中传递。
	*/
}
