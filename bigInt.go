package main

import (
	"fmt"
	"math/big"
	"time"
)

const LIM = 100

var fibs [LIM]*big.Int

func main() {
	result := big.NewInt(0)
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacciData(i)
		fmt.Printf("数列第%d位：%d\n", i+1, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行完成，所耗时间为：%d ns\n", delta)
	fmt.Println(time.Now())
}

func fibonacciData(n int) (res *big.Int) {
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(fibs[n-1], fibs[n-2])
	}
	fibs[n] = res
	return
}
