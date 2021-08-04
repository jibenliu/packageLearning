package main

import (
	"fmt"
	"time"
)

type Pool struct {
	work chan func()   //任务
	sem  chan struct{} //数量
}

func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

func (p *Pool) worker(task func()) {
	defer func() { <-p.sem }()
	for {
		task()
		task = <-p.work
	}
}

func main() {
	pool := New(2) //设置协程池数为 2
	for i := 1; i < 5; i++ { //开启四个任务
		pool.NewTask(func() {
			time.Sleep(2 * time.Second)
			fmt.Println(time.Now())
		})
	}

	time.Sleep(5 * time.Second)
}
