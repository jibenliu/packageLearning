package lockless

import (
	"sync"
	"testing"
)

func TestName(t *testing.T) {

	lq := NewQueue(10)
	w := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func(gi int) {
			lq.PushBack(gi)
			w.Done()
		}(i)
	}

	go func() {
		for {
			lq.PopFront()
			//	time.Sleep(1 * time.Second)
		}
	}()
	w.Wait()
}

var ch = make(chan interface{}, 50000)

func BenchmarkGo_Chan(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- 123
		go func() {
			<-ch
		}()

	}
}

func BenchmarkGo_LockFree(b *testing.B) {
	lq := NewQueue(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lq.PushBack(123)
		go func() {
			lq.PopFront()
		}()

	}
}
