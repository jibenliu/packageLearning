package main

import (
	"fmt"
	"math/big"
	"strconv"
	"unsafe"
)

// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return s
}

func main() {
	fmt.Println(convertToBin(2))
	fmt.Println(convertToBin(19))
	fmt.Println(convertToBin(15))
	fmt.Println(convertToBin(0))
	fmt.Println(convertToBin(12345))
	fmt.Println("-1 << 31 :", -1<<31)
	fmt.Println(unsafe.Sizeof(-1)) //8位 10000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001
	// -1 左移31位 1 00000000 00000000 00000000 00000001 00000000 00000000 00000000 0000000
	// 丢弃最高位，0补最低位 符号位不能丢弃
	fmt.Println("1 << 32 :", 1<<32)
	fmt.Println("1 << 33 :", 1<<33)
	fmt.Println("1 << 35 :", 1<<35)
	fmt.Println("1 << 62 :", 1<<62)
	//fmt.Println("1 << 63 :", 1<<63) 会溢出
	fmt.Println("-1 << 62 :", -1<<62)

	var a uint64
	var b big.Int
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
}
