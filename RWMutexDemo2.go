package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)
	go writeLock2(1)
	go writeLock2(2)
	go writeLock2(3)
	time.Sleep(time.Second * 5)
}
func readLock2(i int) {
	println(i, "read start")
	m.RLock()
	println(i, "reading")
	time.Sleep(time.Second)
	m.RUnlock()
	println(i, "read over")
}
func writeLock2(i int) {
	println(i, "write lock")
	m.Lock()
	println(i, "writing")
	time.Sleep(time.Second)
	m.Unlock()
	println(i, "write over")
}
