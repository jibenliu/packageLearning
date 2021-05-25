package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Printf("你好，世界\n")
		mu.Unlock()
	}()
	mu.Lock()
}
