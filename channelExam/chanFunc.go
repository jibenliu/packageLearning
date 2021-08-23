package main

import "fmt"

type ChannelCounter struct {
	ch     chan func()
	number uint64
}

func NewChannelCounter() *ChannelCounter {
	counter := &ChannelCounter{make(chan func(), 100), 0}
	go func(counter *ChannelCounter) {
		for f := range counter.ch {
			f()
		}
	}(counter)
	return counter
}

func (c *ChannelCounter) Add(num uint64) {
	c.ch <- func() {
		c.number = c.number + num
	}
}

func (c *ChannelCounter) Read() uint64 {
	ret := make(chan uint64)
	c.ch <- func() {
		ret <- c.number
		close(ret)
	}
	return <-ret
}

func main() {
	counter := NewChannelCounter()
	for i := 0; i < 100; i++ {
		counter.Add(uint64(i))
	}

	fmt.Println(counter.Read())
}
