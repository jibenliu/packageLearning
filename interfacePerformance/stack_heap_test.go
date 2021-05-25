package interfacePerformance

import (
	"fmt"
	"strconv"
)

// 堆栈中间内联 mid-stack

type Rectangle struct {
}

//go:noinline
func (r *Rectangle) Height() int {
	h, _ := strconv.ParseInt("7", 10, 0)
	return int(h)
}

func (r *Rectangle) Width() int {
	return 6
}

func (r *Rectangle) Area() int {
	return r.Height() * r.Width()
}

func main() {
	var r Rectangle
	fmt.Println(r.Area())
}

// go build -gcflags='-m=2' stack_heap_test.go 差看一下内联的情况,只支持main包
// go build -gcflags "-m -l"
