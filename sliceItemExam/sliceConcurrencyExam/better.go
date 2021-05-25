package main

import (
	"fmt"
	"sync"
)

type ServiceData struct {
	ch   chan int // 用来 同步的channel
	data []int    // 存储数据的slice
}

func (s *ServiceData) Schedule() {
	for i := range s.ch {
		s.data = append(s.data, i)
	}
}

func (s *ServiceData) Close() {
	close(s.ch)
}

func (s *ServiceData) AddData(v int) {
	s.ch <- v
}
func NewScheduleJob(size int, done func()) *ServiceData {
	s := &ServiceData{
		ch:   make(chan int, size),
		data: make([]int, 0),
	}
	go func() {
		s.Schedule()
		done()
	}()
	return s
}

func main() {
	var (
		wg sync.WaitGroup
		n  = 1000
	)

	c := make(chan struct{})

	s := NewScheduleJob(n, func() {
		c <- struct{}{}
	})

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(v int) {
			defer wg.Done()
			s.AddData(v)
		}(i)
	}

	wg.Wait()
	s.Close()
	<-c
	fmt.Println(len(s.data))
}
