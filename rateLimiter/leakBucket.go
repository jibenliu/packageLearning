package main

import (
	"math"
	"sync"
	"time"
)

// 漏桶
// 一个固定大小的桶，请求按照固定的速率流出
// 如果桶是空的，不需要流出请求
// 请求数大于桶的容量，则抛弃多余请求

type LeakyBucket struct {
	rate       float64    // 每秒固定流出速率
	capacity   float64    // 桶的容量
	water      float64    // 当前桶中请求量
	lastLeakMs int64      // 桶上次漏水微秒数
	lock       sync.Mutex // 锁
}

func (leaky *LeakyBucket) Allow() bool {
	leaky.lock.Lock()
	defer leaky.lock.Unlock()

	now := time.Now().UnixNano() / 1e6
	// 计算剩余水量,两次执行时间中需要漏掉的水
	leakyWater := leaky.water - (float64(now-leaky.lastLeakMs) * leaky.rate / 1000)
	leaky.water = math.Max(0, leakyWater)
	leaky.lastLeakMs = now
	if leaky.water+1 <= leaky.capacity {
		leaky.water++
		return true
	} else {
		return false
	}
}

func (leaky *LeakyBucket) Set(rate, capacity float64) {
	leaky.rate = rate
	leaky.capacity = capacity
	leaky.water = 0
	leaky.lastLeakMs = time.Now().UnixNano() / 1e6
}
