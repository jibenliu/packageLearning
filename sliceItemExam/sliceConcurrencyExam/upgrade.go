package main

import (
	"fmt"
	"sync"
)

func main() {
	slc := make([]int, 0, 1000)
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			slc = append(slc, a)
		}(i)
	}
	wg.Wait()
	fmt.Println(len(slc))

}
