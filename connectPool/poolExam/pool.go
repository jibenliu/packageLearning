package main

import (
	customPool "connectPool/poolExam/src"
	"context"
)

// FROM https://juejin.im/post/5d7471ab51882531ea0649e6


//不释放资源
//func main() {
//	ctx := context.Background()
//	config := &custom_pool.Config{
//		MaxConn: 2,
//		MaxIdle: 1,
//	}
//	conn := custom_pool.Prepare(ctx, config)
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//}

//释放资源
//func main()  {
//	ctx := context.Background()
//	config := &custom_pool.Config{
//		MaxConn: 2,
//		MaxIdle: 1,
//	}
//	conn := custom_pool.Prepare(ctx, config)
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//	go conn.Release(ctx)
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//	if _, err := conn.New(ctx); err != nil {
//		return
//	}
//}

//使用连接池
func main() {

	ctx := context.Background()
	config := &customPool.Config{
		MaxConn: 2,
		MaxIdle: 1,
	}
	conn := customPool.Prepare(ctx, config)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
}
