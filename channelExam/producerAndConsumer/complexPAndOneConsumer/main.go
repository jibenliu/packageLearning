package main

import (
	"fmt"
	"sync"
)

// Msg 生产者共用channel
type Msg struct {
	in int
	ch chan int //双向通信chan
}

// 生产者
func producer(sendChan chan Msg, wg *sync.WaitGroup) {
	recvCh := make(chan int)
	go func() {
		for v := range recvCh {
			fmt.Println("recv ", v)
		}
	}()

	for i := 0; i < 10; i++ {
		sendChan <- Msg{in: i, ch: recvCh}
	}

	wg.Done()
}

// 消费者
func consumer(sendChan chan Msg) {
	for v := range sendChan {
		process(v)
	}
}

// 消息处理函数
func process(msg Msg) {
	msg.ch <- 2 * msg.in
}

func mpsc() {
	// 生产者个数
	pNum := 3
	sendChan := make(chan Msg, 10)

	wg := sync.WaitGroup{}
	wg.Add(pNum)
	for p := 0; p < pNum; p++ {
		go producer(sendChan, &wg)
	}

	// 等待生产者都完成后关闭 sendChan 通知到 消费者
	go func() {
		wg.Wait()
		close(sendChan)
	}()

	consumer(sendChan)
}

func main() {
	mpsc()
}
