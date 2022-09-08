refer:[用 Redis 做一个可靠的延迟队列](https://mp.weixin.qq.com/s/3Gy7-rAxYtCzchqESYGTRw)
## Example

```go
package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/hdt3213/delayqueue"
	"strconv"
	"time"
)

func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	queue := delayqueue.NewQueue("example", redisCli, func(payload string) bool {
		// callback returns true to confirm successful consumption.
		// If callback returns false or not return within maxConsumeDuration, DelayQueue will re-deliver this message
		return true
	})
	// send delay message
	for i := 0; i < 10; i++ {
		err := queue.SendDelayMsg(strconv.Itoa(i), time.Hour, delayqueue.WithRetryCount(3))
		if err != nil {
			panic(err)
		}
	}
	// send schedule message
	for i := 0; i < 10; i++ {
		err := queue.SendScheduleMsg(strconv.Itoa(i), time.Now().Add(time.Hour))
		if err != nil {
			panic(err)
		}
	}
	// start consume
	done := queue.StartConsume()
	<-done
}
```

## options

```
WithLogger(logger *log.Logger)
```

WithLogger customizes logger for queue

```
WithFetchInterval(d time.Duration)
```

WithFetchInterval customizes the interval at which consumer fetch message from redis

```
WithMaxConsumeDuration(d time.Duration)
```

WithMaxConsumeDuration customizes max consume duration

If no acknowledge received within WithMaxConsumeDuration after message delivery, DelayQueue will try to deliver this
message again

```
WithFetchLimit(limit uint)
```

WithFetchLimit limits the max number of messages at one time


```
WithDefaultRetryCount(count uint)
```

WithDefaultRetryCount customizes the max number of retry, it effects of messages in this queue

use WithRetryCount during DelayQueue.SendScheduleMsg or DelayQueue.SendDelayMsg to specific retry count of particular message