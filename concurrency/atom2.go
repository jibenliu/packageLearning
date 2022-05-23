package concurrency

import (
	"sync/atomic"
	"time"
)

type Atomic2Counter struct {
	number uint64
}

func NewAtomic2Counter() *Atomic2Counter {
	return &Atomic2Counter{0}
}

func (c *Atomic2Counter) Add(num uint64) {
	for {
		time.Sleep(time.Second)
		v := atomic.LoadUint64(&c.number)
		if atomic.CompareAndSwapUint64(&c.number, v, v+num) {
			return
		}
	}
}

func (c *Atomic2Counter) Read() uint64 {
	return atomic.LoadUint64(&c.number)
}
