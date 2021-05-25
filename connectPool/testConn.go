package main

import (
	"context"
)

func main() {

	ctx := context.Background()
	config := &Config{
		MaxConn: 2,
		MaxIdle: 1,
	}
	conn := Prepare(ctx, config)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}
	go conn.Release(ctx)
	if _, err := conn.New(ctx); err != nil {
		return
	}

}

