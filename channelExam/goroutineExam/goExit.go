package main

import (
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		runtime.Goexit() //ιεΊεη¨
		println("never executed")
	}()

	wg.Wait()
	println("executed")
}
