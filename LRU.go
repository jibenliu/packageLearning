package main

import (
	"fmt"
	"sync"
)
//refer:https://mp.weixin.qq.com/s/DWD25p9kBkm-nSqDKJErzg
// 1. 固定容量的缓存，当元素个数满了之后，往里面再放元素时会先删除最近最少使用的元素。
// 2. 像map一样，快速存取元素

type Entry struct {
	Key   string
	Value interface{}
	pre   *Entry
	next  *Entry
}

type Cache struct {
	cache    map[string]*Entry
	capacity int
	head     *Entry
	tail     *Entry
}

func NewCache(cap int) *Cache {
	return &Cache{cache: make(map[string]*Entry), capacity: cap}
}

var lock sync.RWMutex

// 把元素放入缓存中，如果缓存满了，则删除最近最少使用的那个元素并返回，把新元素放入缓存中。
// 如果缓存没满，把新元素放入缓存中并返回nil。
func (cache *Cache) Put(key string, val interface{}) interface{} {
	lock.Lock()
	defer lock.Unlock()

	if existVal, exist := cache.cache[key]; exist {
		cache.moveToHead(existVal)
		return nil
	}

	e := &Entry{Key: key, Value: val, next: cache.head}
	if cache.head != nil {
		cache.head.pre = e
	}
	cache.head = e
	if cache.tail == nil {
		cache.tail = e
	}
	cache.cache[key] = e

	if len(cache.cache) <= cache.capacity {
		return nil
	}

	removedEntry := cache.tail
	cache.tail = cache.tail.pre
	removedEntry.pre = nil
	cache.tail.next = nil
	delete(cache.cache, removedEntry.Key)
	return removedEntry.Value
}

// 从缓存中获取元素
func (cache *Cache) Get(key string) interface{} {
	lock.Lock()
	defer lock.Unlock()

	if existVal, exist := cache.cache[key]; exist {
		// 把该元素提到队列头部
		cache.moveToHead(existVal)
		return existVal.Value
	}
	return nil
}

// 把元素提到队列头部
func (cache *Cache) moveToHead(e *Entry) {
	// 元素存在，下面把元素放到最前面去
	if e == cache.head {
		return
	}
	// 元素不是head，那么e.pre一定不为nil
	e.pre.next = e.next
	if e == cache.tail {
		cache.tail = e.pre
	} else {
		e.next.pre = e.pre
	}
	e.pre = nil
	e.next = cache.head
	cache.head.pre = e
	cache.head = e
}

func main()  {
	cache := NewCache(2)
	cache.Put("1", "one")  // 放入元素one，此时one在队列头部
	fmt.Println(cache.Get("1"))  // 此处输出“one”，此时one在队列头部
	cache.Put("2", "two")  // 放入元素two，此时two在队列头部
	fmt.Println(cache.Get("1"))  // 此处输出“one”，此时one在队列头部
	cache.Put("3", "three")  // 放入元素three，总元素个数为3，因此最近最少使用的元素“2”会被删除
	fmt.Println(cache.Get("2"))  // 此处输出nil
	fmt.Println(cache.Get("3"))  // 此处输出“three”
	fmt.Println(cache.Get("3"))  // 此处输出“three”
}