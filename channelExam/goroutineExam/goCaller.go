package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var skip int
		for {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				break
			}
			fmt.Printf("%s:%d\n", file, line)
			skip++
		}
		wg.Done()
	}()

	wg.Wait()
}


// go tool compile -S main.go 来生成对应的汇编代码
// 通过命令 GOSSAFUNC=printList Go run main.go 来生成 SSA 代码看编译过程
// GOSSAFUNC=printList Go run -gcflags="-d=ssa/prove/debug=3" main.go 可以把 pass 背后的逻辑打印出来，它也会生成 SSA 文件来帮助你 debug