package main

import (
	"context"
	"fmt"
	"time"
	"workerPool/wpool"
)

func main() {
	workerCount := 2
	wp := wpool.NewPool(workerCount)

	ctx, cancel := context.WithTimeout(context.TODO(), 6*time.Second)
	defer cancel()

	go wp.GenerateFrom(wpool.GenJobs())

	go wp.Run(ctx)

	for {
		select {
		case r, ok := <-wp.Results():
			if !ok {
				continue
			}
			if r.Err != nil {
				fmt.Printf("job output error :%v\n", r.Err)
				continue
			}
			fmt.Printf("ID is:%d\n", r.Descriptor.ID)
		case <-wp.Done:
			return
		default:
		}
	}
}
