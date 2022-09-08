package main

import (
	"context"
)

func main() {
	InitRedis()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	InitQueue(ctx)
	select {}
}
