package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		var j = i
		wg.Add(1)
		go func() {
			fmt.Printf("你好，世界+%d\n", j)
			wg.Done()
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()
}
