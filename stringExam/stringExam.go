package main

import (
	"fmt"
	"unicode/utf8"
)

// 字符串转字节切片和字符转字节的结果保持一致
func main() {
	for _, v := range []rune("新世界杂货铺") {
		fmt.Printf("%x ", v)
	}
	fmt.Println()
	bs := []byte("新世界杂货铺")
	for len(bs) > 0 {
		r, w := utf8.DecodeRune(bs)
		fmt.Printf("%x ", r)
		bs = bs[w:]
	}
	fmt.Println()
}
