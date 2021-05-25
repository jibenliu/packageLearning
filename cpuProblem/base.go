package main

import (
	"fmt"
)

type ZeroEvenOdd1 struct {
	n                int
	streamEvenToZero chan struct{}
	streamOddToZero  chan struct{}
	streamZeroToEven chan struct{}
	streamZeroToOdd  chan struct{}
	streamZeroToEnd  chan struct{}
}

func (z *ZeroEvenOdd1) Zero(printNumber func(int)) {
	for i := 0; i < z.n; i++ {
		select {
		case <-z.streamOddToZero:
			printNumber(0)
			z.streamZeroToEven <- struct{}{}
		case <-z.streamEvenToZero:
			printNumber(0)
			z.streamZeroToOdd <- struct{}{}
		default:
			i--
		}

		if 0 == z.n%2 {
			<-z.streamEvenToZero //等待 Even() 結束，自己再結束
		} else {
			<-z.streamOddToZero //等待 Odd() 結束，自己再結束
		}

		z.streamZeroToEnd <- struct{}{}
	}
}

func (z *ZeroEvenOdd1) Even(printNumber func(int)) {
	evenUpper := z.n - z.n%2
	// fmt.Println("evenUpper:", evenUpper)
	for i := 2; i <= evenUpper; {
		<-z.streamZeroToEven
		printNumber(i)
		i += 2
		z.streamEvenToZero <- struct{}{}
	}
}

func (z *ZeroEvenOdd1) Odd(printNumber func(int)) {
	oddUpper := ((z.n + 1) - (z.n+1)%2) - 1
	for i := 1; i <= oddUpper; i += 2 {
		<-z.streamZeroToOdd
		printNumber(i)
		z.streamOddToZero <- struct{}{}
	}
}

func PrintNumber1(x int) {
	fmt.Printf("%d", x)
}

func main() {
	var PrintZeroEvenOdd = func(testNum int) {
		var zeo = &ZeroEvenOdd1{
			n:                testNum,
			streamEvenToZero: make(chan struct{}),
			streamOddToZero:  make(chan struct{}),
			streamZeroToEven: make(chan struct{}),
			streamZeroToOdd:  make(chan struct{}),
			streamZeroToEnd:  make(chan struct{}),
		}

		go func() { zeo.streamEvenToZero <- struct{}{} }() //給起頭的火種
		go zeo.Zero(PrintNumber)
		go zeo.Even(PrintNumber)
		go zeo.Odd(PrintNumber)
		<-zeo.streamZeroToEnd //等待 Zero() 送出結束訊號
		fmt.Println()
	}

	for testNum := range [14]int{} {
		fmt.Printf("Case %2d: ", testNum)
		PrintZeroEvenOdd(testNum)
	}
}
