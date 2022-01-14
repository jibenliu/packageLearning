package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			start := time.Now()
			query()
			elapsed := time.Since(start)
			fmt.Printf("goroutine id %d costs time %s\n", i, elapsed)
		}(i)
	}
	select {}
}

func query() {
	once.Do(func() {
		time.Sleep(5 * time.Second)
	})
}
