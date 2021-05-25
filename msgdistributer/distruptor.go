package main

import (
	"fmt"
	"github.com/smartystreets-prototypes/go-disruptor"
	"log"
	"time"
)

func main() {
	iterations := int64(1000 * 1000 * 20)

	consumer := SampleConsumer{}
	myDisruptor := build(consumer)

	go func() {
		var sequence int64 = -1
		for sequence < iterations {
			sequence = myDisruptor.Reserve(ReserveMany)
			for i := sequence - (ReserveMany - 1); i <= sequence; i++ {
				ringBuffer[i&RingBufferMask] = i
			}
			myDisruptor.Commit(sequence, sequence)
		}

		_ = myDisruptor.Close()
	}()
	t := time.Now().UnixNano()
	myDisruptor.Read()
	t = (time.Now().UnixNano() - t) / 1000000 //ms
	fmt.Printf("opsPerSecond: %d\n", iterations*1000/t)
}

type SampleConsumer struct{}

func (s SampleConsumer) Consume(lower, upper int64) {
	var message int64
	for lower <= upper {
		message = ringBuffer[lower&RingBufferMask]
		if message != lower {
			log.Panicf("race condition: Sequence: %d, Message: %d", lower, message)
		}
		lower++
	}
}

func build(consumers ...disruptor.Consumer) disruptor.Disruptor {
	return disruptor.New(
		disruptor.WithCapacity(RingBufferSize),
		disruptor.WithConsumerGroup(consumers...))
}

const (
	RingBufferSize = 1024 * 64
	RingBufferMask = RingBufferSize - 1
	ReserveOne     = 1
	ReserveMany    = 16
)

var ringBuffer = [RingBufferSize]int64{}
