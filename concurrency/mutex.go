package concurrency

import (
	"sync"
	"time"
)

type MutexCounter struct {
	mu     *sync.RWMutex
	number uint64
}

func NewMutexCounter() *MutexCounter {
	return &MutexCounter{&sync.RWMutex{}, 0}
}

func (c *MutexCounter) Add(num uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	time.Sleep(time.Second)
	c.number = c.number + num
}

func (c *MutexCounter) Read() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.number
}
