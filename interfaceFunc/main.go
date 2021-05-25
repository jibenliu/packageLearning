package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(i interface{}) {
	fmt.Println("from struct ", i)
}

type FuncCaller func(interface{}) //用函数实现接口

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func main() {
	var invoker Invoker

	s := new(Struct)
	invoker = s

	invoker.Call("hello")

	invoker = FuncCaller(func(i interface{}) {
		fmt.Println("from function", i)
	})

	invoker.Call("hello")
	a := []int{1, 2, 3}
	fmt.Printf("the variable a address is %p ,\n", a)
}
