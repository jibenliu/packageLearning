package main

import (
	"fmt"
	"time"
)
/**
	协程的执行顺序大致如下所示：
		从寄存器读取 a 的值；
		然后做加法运算；
		最后写到寄存器。
	假如有一个协程取得 a 的值为 3，然后执行加法运算，此时又有一个协程对 a 进行取值，得到的值同样是 3，最终两个协程的返回结果是相同的
 */
func main() {
	var a = 0

	for i := 0; i < 1000; i++ {
		go func(idx int) {
			a += 1
			fmt.Println(a)
		}(i)
	}
	time.Sleep(time.Second)
}
