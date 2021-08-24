package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func doAPICall(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	req, err := http.NewRequest("GET", "https://httpstat.us/200", nil)
	if err != nil {
		return err
	}

	// The httpstat.us API accepts a sleep parameter which sleeps the request for the
	// passed time in ms
	q := req.URL.Query()
	sleepMin := 1000
	sleepMax := 4000
	q.Set("sleep", fmt.Sprintf("%d", rand.Intn(sleepMax-sleepMin)+sleepMin))
	req.URL.RawQuery = q.Encode()

	// Make the request to the API in an anonymous function, using a channel to
	// communicate the results
	c := make(chan error, 1)
	go func() {
		// For the purposes of this example, we're not doing anything with the response.
		_, err := http.DefaultClient.Do(req)
		c <- err
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-c:
		return err
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var (
		closing     = make(chan struct{})
		ticker      = time.NewTicker(1 * time.Second)
		logger      = log.New(os.Stderr, "", log.LstdFlags)
		batchSize   = 6
		jobs        = make(chan int, batchSize)
		wg          sync.WaitGroup
		ctx, cancel = context.WithCancel(context.Background())
	)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
		<-signals
		logger.Println("Initiating shutdown of producer.")
		cancel()
		close(closing)
	}()

loop:
	for {
		select {
		case <-closing:
			break loop
		case <-ticker.C:
			for n := 0; n < batchSize; n++ {
				wg.Add(1)
				jobs <- n
				go func() {
					_ = doAPICall(ctx, &wg)
				}()
			}
			wg.Wait()
			logger.Printf("Completed doing %d things.", batchSize)
		}
	}
}
