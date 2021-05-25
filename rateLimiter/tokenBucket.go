package main

import (
	"sync"
	"time"
)

type TokenBucket struct {
	rate         int64 //固定的token放入速率, r/s
	capacity     int64 //桶的容量
	tokens       int64 //桶中当前token数量
	lastTokenSec int64 //上次向桶中放令牌的时间的时间戳，单位为秒

	lock sync.Mutex
}

func (bucket *TokenBucket) Take() bool {
	bucket.lock.Lock()
	defer bucket.lock.Unlock()

	now := time.Now().Unix()
	bucket.tokens = bucket.tokens + (now-bucket.lastTokenSec)*bucket.rate // 先添加令牌
	if bucket.tokens > bucket.capacity {
		bucket.tokens = bucket.capacity
	}
	bucket.lastTokenSec = now
	if bucket.tokens > 0 {
		// 还有令牌，领取令牌
		bucket.tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}

func (bucket *TokenBucket) Init(rate, cap int64) {
	bucket.rate = rate
	bucket.capacity = cap
	bucket.tokens = 0
	bucket.lastTokenSec = time.Now().Unix()
}
