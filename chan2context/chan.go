package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func doSomething(ch chan int) {
	fmt.Printf("Received job %d\n", <-ch)
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var (
		closing   = make(chan struct{})
		ticker    = time.NewTicker(1 * time.Second)
		logger    = log.New(os.Stderr, "", log.LstdFlags)
		batchSize = 6
		jobs      = make(chan int, batchSize)
	)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
		<-signals
		close(closing)
	}()

loop:
	for {
		select {
		case <-closing:
			break loop
		case <-ticker.C:
			for n := 0; n < batchSize; n++ {
				jobs <- n
				go doSomething(jobs)
			}
			logger.Printf("Completed doing %d things.", batchSize)
		}
	}
}
