package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel 是一个能接收 DataEvent 的 channel
type DataChannel chan DataEvent

// DataChannelSlice 是一个包含 DataChannels 数据的切片
type DataChannelSlice []DataChannel

// EventBus 事件总线 存储有关订阅者感兴趣的特定主题的信息
type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

// Subscribe 订阅主题
func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

// Publish 发布主题
func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		// 这样做是因为切片引用相同的数组，即使它们是按值传递的
		// 因此我们正在使用我们的元素创建一个新切片，从而能正确地保持锁定
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

var eb = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func publishTo(topic string, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)

	go publishTo("topic1", "Hi topic 1")
	go publishTo("topic2", "Welcome to topic 2")

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}
