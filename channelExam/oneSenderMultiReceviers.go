package main

import (
	"fmt"
	"strings"
	"sync"
)

const (
	NumReceivers1 = 20
)

var wg sync.WaitGroup
var strMessages = []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii", "jjj", "kkk"}

func sender(msg []string, ch chan string) {
	for i, v := range msg {
		ch <- v
		if i == len(msg)-1 {
			close(ch)
		}
	}
}

func receiver(ch chan string) {
	defer wg.Done()
	for value := range ch {
		fmt.Println(strings.ToUpper(value))
	}
}

func main() {
	ch := make(chan string, 20)

	go sender(strMessages, ch)

	wg.Add(NumReceivers1)

	for i := 0; i < NumReceivers1; i++ {
		go receiver(ch)
	}
	wg.Wait()
}
