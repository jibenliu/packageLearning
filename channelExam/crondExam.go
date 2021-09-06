package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// refer:https://mp.weixin.qq.com/s/tuVBwiO78CgEgQ6CPHe3sg
func main() {
	useConn()
	useBufferedChan()
	useUnbufferedChan()
}

func useConn() {
	ready := 0
	c := sync.NewCond(&sync.Mutex{})

	// 起10个协程，随机等待后模拟运动员就位，并记录就位人数（加锁更新）
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			// 这里用signal也可以，因为等待者只有一个main goroutine
			c.Broadcast()
		}(i)
	}

	// 等待条件满足：10人都就位
	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	// 所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}

func useUnbufferedChan() {
	// unbuffered
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- 1
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
	println("all is ready!")
}

func useBufferedChan() {
	// buffered
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	for len(ch) != 10 {
		time.Sleep(time.Millisecond)
	}
	println("all is ready!")
}
