package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GoRun struct {
	ch chan struct{}
}

func NewGoRun(maxRun int) *GoRun {
	return &GoRun{ch: make(chan struct{}, maxRun)}
}

func (p *GoRun) Done() {
	<-p.ch
}

func (p *GoRun) Count() int {
	return len(p.ch)
}

func (p *GoRun) Max() int {
	return cap(p.ch)
}

func (p *GoRun) Run(f func()) {
	select {
	case p.ch <- struct{}{}:
		go func(f func()) {
			defer p.Done()
			f()
		}(f)
	default:
		return
	}
}

func T() {
	p := NewGoRun(100)
	for {
		p.Run(func() {
			rand.Seed(time.Now().UnixNano())
			x := rand.Intn(p.Max() * 100)
			time.Sleep(time.Millisecond * time.Duration(x))
		})
		fmt.Println(p.Count())
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go T()
	select {}
}
