package concurrency

import (
	"sync/atomic"
	"time"
)

type AtomicCounter struct {
	number uint64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{0}
}

func (c *AtomicCounter) Add(num uint64) {
	time.Sleep(time.Second)
	atomic.AddUint64(&c.number, num)
}

func (c *AtomicCounter) Read() uint64 {
	return atomic.LoadUint64(&c.number)
}
