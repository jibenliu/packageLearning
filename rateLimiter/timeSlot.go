package main

import (
	"sync"
	"time"
)

type timeSlot struct {
	timestamp time.Time // 这个timeSlot的时间起点
	count     int       // 落在这个timeSlot内的请求数
}

// 统计整个时间窗口中已经发生的请求次数
func countReq(win []*timeSlot) int {
	var count int
	for _, ts := range win {
		count += ts.count
	}
	return count
}

type SlidingWindowLimiter struct {
	mu           sync.Mutex    // 互斥锁保护其他字段
	SlotDuration time.Duration // time slot的长度
	WinDuration  time.Duration // sliding window的长度
	numSlots     int           // window内最多有多少个slot
	windows      []*timeSlot
	maxReq       int // 大窗口时间内允许的最大请求数
}

func NewSliding(slotDuration time.Duration, winDuration time.Duration, maxReq int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		SlotDuration: slotDuration,
		WinDuration:  winDuration,
		numSlots:     int(winDuration / slotDuration),
		maxReq:       maxReq,
	}
}

func (l *SlidingWindowLimiter) validate() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	// 已经过期的time slot移出时间窗
	timeoutOffset := -1
	for i, ts := range l.windows {
		if ts.timestamp.Add(l.WinDuration).After(now) {
			break
		}
		timeoutOffset = i
	}
	if timeoutOffset > -1 {
		l.windows = l.windows[timeoutOffset+1:]
	}

	// 判断请求是否超限
	var result bool
	if countReq(l.windows) < l.maxReq {
		result = true
	}

	// 记录这次的请求数
	var lastSlot *timeSlot
	if len(l.windows) > 0 {
		lastSlot = l.windows[len(l.windows)-1]
		if lastSlot.timestamp.Add(l.SlotDuration).Before(now) {
			// 如果当前时间已经超过这个时间插槽的跨度，那么新建一个时间插槽
			lastSlot = &timeSlot{timestamp: now, count: 1}
			l.windows = append(l.windows, lastSlot)
		} else {
			lastSlot.count++
		}
	} else {
		lastSlot = &timeSlot{timestamp: now, count: 1}
		l.windows = append(l.windows, lastSlot)
	}
	return result
}
