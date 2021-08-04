package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	NumPublishers := runtime.NumCPU()
	totalIterations := int64(1000 * 1000 * 20)
	iterations := totalIterations / int64(NumPublishers)
	totalIterations = iterations * int64(NumPublishers)
	channel := make(chan int64, 1024*64)
	var wg sync.WaitGroup
	wg.Add(NumPublishers + 1)
	var readerWG sync.WaitGroup
	readerWG.Add(1)
	for i := 0; i < NumPublishers; i++ {
		go func() {
			wg.Done()
			wg.Wait()
			for i := int64(0); i < iterations; {
				select {
				case channel <- i:
					i++
				default:
					continue
				}
			}
		}()
	}
	go func() {
		for i := int64(0); i < totalIterations; i++ {
			select {
			case msg := <-channel:
				if NumPublishers == 1 && msg != i {
					//panic("Out of sequence")
				}
			default:
				continue
			}
		}
		readerWG.Done()
	}()
	wg.Done()
	t := time.Now().UnixNano()
	wg.Wait()
	readerWG.Wait()
	t = (time.Now().UnixNano() - t) / 1000000 //ms
	fmt.Printf("opsPerSecond: %d\n", totalIterations*1000/t)
}
