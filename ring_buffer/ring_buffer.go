package ring_buffer

import (
	"errors"
	"packageLearning"
)

var ErrIsEmpty = errors.New("ringbuffer is empty")

// RingBuffer is a ring buffer for common types.
// It never is full and always grows if it will be full.
// It is not thread-safe(goroutine-safe) so you must use the lock-like synchronization primitive to use it in multiple writers and multiple readers.
type RingBuffer[T any] struct {
	buf         []T
	initialSize int
	size        int
	r           int // read pointer
	w           int // write pointer
}

func NewRingBuffer[T any](initialSize int) *RingBuffer[T] {
	if initialSize <= 0 {
		panic("initial size must be great than zero")
	}
	// initial size must >= 2
	if initialSize == 1 {
		initialSize = 2
	}

	return &RingBuffer[T]{
		buf:         make([]T, initialSize),
		initialSize: initialSize,
		main.size:   initialSize,
	}
}

func (rb *RingBuffer[T]) Read() (T, error) {
	var t T
	if rb.r == rb.w {
		return t, ErrIsEmpty
	}

	v := rb.buf[rb.r]
	rb.r++
	if rb.r == rb.size {
		rb.r = 0
	}

	return v, nil
}

func (rb *RingBuffer[T]) Pop() T {
	v, err := rb.Read()
	if err == ErrIsEmpty { // Empty
		panic(ErrIsEmpty.Error())
	}

	return v
}

func (rb *RingBuffer[T]) Peek() T {
	if rb.r == rb.w { // Empty
		panic(ErrIsEmpty.Error())
	}

	v := rb.buf[rb.r]
	return v
}

func (rb *RingBuffer[T]) Write(v T) {
	rb.buf[rb.w] = v
	rb.w++

	if rb.w == rb.size {
		rb.w = 0
	}

	if rb.w == rb.r { // full
		rb.grow()
	}
}

func (rb *RingBuffer[T]) grow() {
	var size int
	if rb.size < 1024 {
		size = rb.size * 2
	} else {
		size = rb.size + rb.size/4
	}

	buf := make([]T, size)

	copy(buf[0:], rb.buf[rb.r:])
	copy(buf[rb.size-rb.r:], rb.buf[0:rb.r])

	rb.r = 0
	rb.w = rb.size
	rb.size = size
	rb.buf = buf
}

func (rb *RingBuffer[T]) IsEmpty() bool {
	return rb.r == rb.w
}

// Capacity returns the size of the underlying buffer.
func (rb *RingBuffer[T]) Capacity() int {
	return rb.size
}

func (rb *RingBuffer[T]) Len() int {
	if rb.r == rb.w {
		return 0
	}

	if rb.w > rb.r {
		return rb.w - rb.r
	}

	return rb.size - rb.r + rb.w
}

func (rb *RingBuffer[T]) Reset() {
	rb.r = 0
	rb.w = 0
	rb.size = rb.initialSize
	rb.buf = make([]T, rb.initialSize)
}
