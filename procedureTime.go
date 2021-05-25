package main

import (
	"fmt"
	"time"
)

//获取流程运行时间
func main() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)

	elapsed2 := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时2：", elapsed2)
}
