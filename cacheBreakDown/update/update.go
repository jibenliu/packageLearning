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
	once *sync.Once
}

type OrderSever struct {
	cache map[string]*CacheEntry
	mutex sync.Mutex
}

func (order *OrderSever) Query(key string) ([]byte, error) {
	order.mutex.Lock()
	if order.cache == nil {
		order.cache = make(map[string]*CacheEntry)
	}
	entry, ok := order.cache[key]

	// 找不到就初始化一个
	if !ok {
		entry = &CacheEntry{
			data: make([]byte, 0),
			err:  nil,
			once: new(sync.Once),
		}
		order.cache[key] = entry
	}
	order.mutex.Unlock()

	// 我只执行一次。
	entry.once.Do(func() {
		entry.data, entry.err = getOrder()
	})
	//数据被赋值了，返回
	return entry.data, entry.err
}

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
			data, _ := server.Query("goroutine")
			fmt.Printf("goroutine id %d get data %s\n", id, string(data))
		}(id)
		id++
	}
}
