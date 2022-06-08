package main

import (
	"context"
	"fmt"
	"io"
	"time"
)

const defaultPingInterval = 5

// Pinger 发起ping请求
// 1.需要一个goroutine定期到发送ping消息
// 2.如果最近收到远程服务的回复，就不需要发送不必要的ping消息，因此需要可以重置ping计时器功能
func Pinger(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	var interval time.Duration
	select {
	case <-ctx.Done():
		return
	case interval = <-reset: //读取更新的心跳间隔时间
	default:
	}
	if interval < 0 {
		interval = defaultPingInterval
	}
	timer := time.NewTimer(interval)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case newInterval := <-reset:
			if !timer.Stop() {
				<-timer.C
			}
			if newInterval > 0 {
				interval = newInterval
			}
		case <-timer.C:
			if _, err := w.Write([]byte("ping")); err != nil {
				//在此跟踪并执行连续超时
				return
			}
		}
		_ = timer.Reset(interval) //重制心跳上报时间间隔
	}
}

// ExamplePinger
// @Description: 模拟ping请求
func ExamplePinger() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	r, w := io.Pipe() //代替网络连接net.Conn
	done := make(chan struct{})
	resetTimer := make(chan time.Duration, 1)
	resetTimer <- time.Second //ping间隔初始值

	go func() {
		Pinger(ctx, w, resetTimer)
		close(done)
	}()
	receivePing := func(d time.Duration, r io.Reader) {
		if d >= 0 {
			fmt.Printf("resetting time (%s)\n", d)
			resetTimer <- d
		}

		now := time.Now()
		buf := make([]byte, 1024)
		n, err := r.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("received %q (%s)\n", buf[:n], time.Since(now).Round(100*time.Millisecond))
	}
	for i, v := range []int64{0, 200, 300, 0, -1, -1, -1} {
		fmt.Printf("Run %d\n", i+1)
		receivePing(time.Duration(v)*time.Millisecond, r)
	}
	cancelFunc() //取消context使pinger退出
	<-done
}

func main() {
	ExamplePinger()
}
