package delayQueue

import "time"

// Job 任务
type Job struct {
	Id         string `msgpack:"1"` // 任务id
	Topic      string `msgpack:"2"` // 消息名
	Delay      int64  `msgpack:"3"` // 延迟时间
	PlayLoad   []byte `msgpack:"4"` // 消息体
	Timestamp  int64  `msgpack:"5"` // 消息投递时间
	TryCount   int32  `msgpack:"6"` // 重试次数 最大丢弃
	TryTimeOut int64  `msgpack:"7"` // 超时重试事件
}

// Dispatch 分发普通任务
func (c *QueueClient) Dispatch(topic string, payload []byte) error {
	return c.queue.Push(&Job{
		Topic:     topic,
		PlayLoad:  payload,
		Delay:     0,
		Timestamp: time.Now().Unix(),
	})
}

// DispatchDelaySecond 分发延迟任务
func (c *QueueClient) DispatchDelaySecond(topic string, payload []byte, delaySec int) error {
	return c.queue.DelayJob(&Job{
		Topic:     topic,
		PlayLoad:  payload,
		Delay:     int64(delaySec),
		Timestamp: time.Now().Unix(),
	})
}
