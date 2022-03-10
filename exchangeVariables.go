package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a = 5
	var b = 10
	fmt.Printf("原始a的内存地址是：%p\n", &a)
	//var iPtr = uintptr(a)
	var iPtr = rune(a)
	fmt.Printf("原始a的指针指向的内存地址是：%p %v\n", &iPtr, iPtr)
	var bi = &b
	fmt.Printf("原始b的内存地址是：%p\n", &b)
	fmt.Printf("原始b的指针指向的内存地址是：%p\n", &bi)
	exchangeVariables(a, b)
	fmt.Println("a = ", a, " b = ", b) // 怎么输出 a = 10 b = 5
}

func exchangeVariables(a, b int) {
	fmt.Printf("函数内a的内存地址是：%p\n", &a)
	var iPtr = (*reflect.StringHeader)(unsafe.Pointer(&a))
	fmt.Printf("函数内a的指针指向的内存地址是：%p\n", &iPtr)
	fmt.Printf("函数内b的内存地址是：%p\n", &b)
	var bi = &b
	fmt.Printf("函数内b的指针指向的内存地址是：%p\n", &bi)
	//reflect.ValueOf(&a).Elem().SetPointer(reflect.ValueOf(&b).Addr())
	//reflect.ValueOf(&b).Elem().SetPointer(*a)
	//fmt.Println(a)
	//fmt.Println(b)
}
