package concurrency

import "time"

type Counter struct {
	number uint64
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Add(num uint64) {
	time.Sleep(time.Second) //模拟耗时操作
	c.number = c.number + num
}

func (c *Counter) Read() uint64 {
	return c.number
}
