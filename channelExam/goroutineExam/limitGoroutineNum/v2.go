package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type GoRunV2 struct {
	max   int
	count int
	sync.Mutex
}

func NewGoRunV2(max int) *GoRunV2 {
	return &GoRunV2{max: max}
}

func (p *GoRunV2) CanRun() bool {
	p.Lock()
	defer p.Unlock()
	if p.count < p.max {
		p.count++
		return true
	}
	return false
}
func (p *GoRunV2) Done() {
	p.Lock()
	defer p.Unlock()
	p.count--
}

func (p *GoRunV2) Count() int {
	return p.count
}

func (p *GoRunV2) Run(f func()) {
	if !p.CanRun() {
		return
	}
	go func(f func()) {
		defer p.Done()
		f()
	}(f)
}

func TV2(){
	p := NewGoRunV2(100)
	for {
		p.Run(func() {
			rand.Seed(time.Now().UnixNano())
			x := rand.Intn(p.max * 100) //生成0-99随机整数
			time.Sleep(time.Millisecond * time.Duration(x))
		})
		fmt.Println(p.Count())
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go TV2()
	select {}
}
