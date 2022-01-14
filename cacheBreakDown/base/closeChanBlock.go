package main

import (
	"fmt"
	"time"
)

var ch = make(chan struct{})

func main() {
	start := time.Now()
	// close 必须在 chan 输出之前，不然chan阻塞导致close无法进入逻辑
	go query()
	<-ch
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func query() {
	time.Sleep(5 * time.Second)
	close(ch)
}
