package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)
	go readLock(1)
	go readLock(2)
	time.Sleep(2 * time.Second)

}

func readLock(i int) {
	println(i, "read start")
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	println(i, "read over")
}
