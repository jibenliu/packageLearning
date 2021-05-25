package main

import (
	"fmt"
	"runtime"
)

type panicContext struct {
	str string
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()
}

func main() {
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		panic(&panicContext{
			"手动触发panic",
		})
	})

	ProtectRun(func() {
		fmt.Println("手动宕机前")
		var p *int
		*p = 1
	})

	fmt.Println("宕机恢复证明")
}
