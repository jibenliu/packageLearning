package main

import "time"

type retryQueue struct {
	buckets       [][]retryItem // each bucket represents a 1 second interval
	currentTime   time.Time
	currentOffset int
}

type retryItem struct {
	id   [20]byte  // ID of the item to retry
	time time.Time // exact time at which the item has to be retried
}

// time.Time 的结构体中包含了一个指针成员 loc 。在 retryItem 中使用它会导致 GC 每次经过堆上的这块区域时。 都要去追踪到结构体里面的指针。
