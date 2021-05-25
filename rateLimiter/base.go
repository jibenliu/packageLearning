package main

import (
	"sync"
	"time"
)

type LimitRate struct {
	rate  int           //阀值
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //收到请求数
	lock  sync.Mutex    //锁
}

func (limit *LimitRate) Allow() bool {
	limit.lock.Lock()
	defer limit.lock.Unlock()

	// 判断收到请求数是否达到阀值
	if limit.count == limit.rate-1 {
		now := time.Now()
		// 达到阀值后，判断是否是请求周期内
		if now.Sub(limit.begin) >= limit.cycle {
			limit.Reset(now)
			return true
		}
		return false
	} else {
		limit.count++
		return true
	}
}

func (limit *LimitRate) Set(rate int, cycle time.Duration) {
	limit.rate = rate
	limit.begin = time.Now()
	limit.cycle = cycle
	limit.count = 0
}

func (limit *LimitRate) Reset(begin time.Time) {
	limit.begin = begin
	limit.count = 0
}
