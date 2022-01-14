package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CacheEntry struct {
	data []byte
	err  error
	wait chan struct{}
}

type OrderSever struct {
	cache map[string]*CacheEntry
	mutex sync.Mutex
}

func (order *OrderSever) Query(key string, id int) ([]byte, error) {
	order.mutex.Lock()
	if order.cache == nil {
		order.cache = make(map[string]*CacheEntry)
	}

	//已经有其他兄弟请求了，你等等
	if entry, ok := order.cache[key]; ok {
		order.mutex.Unlock()

		// 会阻塞，除非其他协程close了chan
		start := time.Now()
		<-entry.wait
		elapsed := time.Since(start)

		fmt.Printf("goroutine id %d release block costs time is %s\n", id, elapsed)
		return entry.data, entry.err
	}
	entry := &CacheEntry{
		data: make([]byte, 0),
		wait: make(chan struct{}),
	}
	order.cache[key] = entry
	order.mutex.Unlock()
	// 请求数据
	entry.data, entry.err = getOrder()
	// 请求数据完毕，通知其他兄弟可以拿数据了。
	close(entry.wait)
	return entry.data, nil
}

//外部服务
func getOrder() ([]byte, error) {
	time.Sleep(10 * time.Second)
	return []byte("hello world"), nil
}

var random = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	server := new(OrderSever)
	var id = 0
	for {
		time.Sleep(time.Duration(random.Intn(2)) * time.Second)
		go func(id int) {
			data, _ := server.Query("goroutine", id)
			fmt.Printf("goroutine id %d get data %s\n", id, string(data))
		}(id)
		id++
	}
}
