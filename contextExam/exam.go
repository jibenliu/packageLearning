package main

import (
	"context"
	"sync"
	"time"
)

type cancelCtx struct {
	context.Context

	// 保护之后的字段
	mu       sync.Mutex
	done     chan struct{}
	children map[canceler]struct{}
	err      error
}

// A canceler is a context type that can be canceled directly. The
// implementations are *cancelCtx and *timerCtx.
type canceler interface {
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}

// &cancelCtxKey is the key that a cancelCtx returns itself for.
var cancelCtxKey int

func (c *cancelCtx) Value(key interface{}) interface{} {
	if key == &cancelCtxKey {
		return c
	}
	return c.Context.Value(key)
}

func (c *cancelCtx) Done() <-chan struct{} {
	c.mu.Lock()
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	c.mu.Unlock()
	return d
}

func (c *cancelCtx) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

func main() {
	ctx := context.Background()
	context.WithValue(ctx, "", "")
	t := time.Now().Add(time.Minute * -1)
	parent, cancel := context.WithDeadline(ctx, t)
	defer cancel()
	ctx, canc := context.WithTimeout(parent, 2*time.Second)
	defer canc()

	cx, c := context.WithCancel(ctx)
	defer c()
}
