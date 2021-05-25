package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	S *string
}

func (f *Foo) String() string {
	return *f.S
}

type FooTrick struct {
	S unsafe.Pointer
}

func (f *FooTrick) String() string {
	return *(*string)(f.S)
}

func NewFoo(s string) Foo {
	return Foo{S: &s}
}

func NewFooTrick(s string) FooTrick {
	return FooTrick{S: noescape(unsafe.Pointer(&s))}
}

//这个函数会把指针参数从编译器的逃逸分析中隐藏掉
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func main() {
	s := "hello"
	f1 := NewFoo(s)
	f2 := NewFooTrick(s)
	s1 := f1.String()
	s2 := f2.String()
	fmt.Println(s1)
	fmt.Println(s2)
}


// go build -gcflags '-m -m' noescape.go  逃逸分析