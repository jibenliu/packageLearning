package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//n := 2000
	n := 2000000000

	/** 并发流程1 **/
	t1 := time.Now()
	sum1 := 0
	for i := 0; i < n; i++ {
		sum1 += i
	}
	expire1 := time.Since(t1)
	fmt.Printf("流程1总耗时：%dms,结果为%d\n", expire1/1e6, sum1)

	//---------------华丽的分割线--------------
	/** 并发流程2 **/
	t2 := time.Now()
	cpuNum := runtime.NumCPU()
	period := n / cpuNum
	sumChan := make(chan int, cpuNum)

	sum2 := 0
	for j := 0; j < cpuNum; j++ {
		go func(j int) {
			temp := 0
			start := period * j
			end := period * (j + 1)
			for k := start; k < end; k++ {
				temp += k
			}
			sumChan <- temp
		}(j)
	}
	for i := 0; i < cpuNum; i++ {
		tmp := <-sumChan
		sum2 += tmp
	}

	expire2 := time.Since(t2)
	fmt.Printf("流程2总并发数%d,总耗时：%dms,结果为%d\n", cpuNum, expire2/1e6, sum2)
}
