package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type ProcessMsg func(ctx context.Context, originMsg interface{}) (resMsg interface{})

var (
	wg3 sync.WaitGroup
)

func senders3(i int, msg []string, stopCh chan struct{}, ch chan string) {
	for _, v := range msg {
		ch <- v
	}
	if i == -1 {
		time.Sleep(time.Millisecond)
		close(stopCh)
	}

}

func receivers3(ctx context.Context, stopCh chan struct{}, ch chan string, process ProcessMsg) {
	defer wg3.Done()

	for {
		select {
		case <-stopCh:
			return

		case value := <-ch:
			resMsg := process(ctx, value)
			fmt.Println(resMsg)
		default:
		}
	}
}

func processMsg(ctx context.Context, originMsg interface{}) (resMsg interface{}) {
	msg := originMsg.(string)
	msg = strings.ToUpper(msg)
	return msg
}

func main() {
	ctx := context.Background()

	ch := make(chan string, 100)
	stopCh := make(chan struct{})

	msg := [][]string{{"abc", "def", "ghi"}, {"123", "234", "345"}, {"a1a", "b2b", "c3c"}}

	length := len(msg)

	for i := 0; i < length; i++ {
		j := i
		if i == length-1 {
			j = -1
		}

		go senders3(j, msg[i], stopCh, ch)
	}

	wg3.Add(length)
	for i := 0; i < length; i++ {
		go receivers3(ctx, stopCh, ch, processMsg)
	}
	wg3.Wait()

}
