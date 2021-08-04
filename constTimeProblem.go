package main

import (
	"fmt"
	"reflect"
	"time"
)

const a = 1

func testConst(a time.Duration) {
	fmt.Printf("%v \n", a)
}
func main() {
	b := a + 1
	// 在 const 修饰的变量如果没有明确指明类型时，会根据需要做类型转换; 而非 const 修饰的变量不会
	fmt.Printf("---- a = %v \n", reflect.TypeOf(a))
	testConst(a)
	fmt.Printf("---- b = %v \n", reflect.TypeOf(b))
	// test(b)  //> 打开注释后, 会编译不通过
}
