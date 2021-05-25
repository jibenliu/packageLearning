package main

import "time"

type retryItem struct {
	id   [20]byte // ID of the item to retry
	nsec uint32
	sec  int64
}

func (item *retryItem) time() time.Time {
	return time.Unix(item.sec, int64(item.nsec))
}

func makeRetryItem(id ksuid.KSUID, time time.Time) retryItem {
	return retryItem{
		id:   id,
		nsec: uint32(time.Nanosecond()),
		sec:  time.Unix(),
	}
}

//现在的 retryItem 不包含任何指针。这大大降低了 gc 压力，因为 retryItem  的整个足迹都可以在编译时知道
