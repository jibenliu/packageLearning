package main

import (
	"fmt"
	"math/rand"
)

type I10 interface {
	int
}

type I11 interface {
	string
}
type S struct {
	b bool
	e error
}

type I12 interface {
	S
}

type I20 interface {
	~int
}

type I21 interface {
	~string
}

/**
近似类型~T必须是底层类型自身, 不能是基于底层类型自定义的新类型
在自定义类型小节中所有自定义的A~Z类型都被限制
近似类型~T不能是接口类型
*/

// 正确的近似类型
type I23 interface {
	~int
	~int8
	~int16
	~rune
	~int32
	~int64
	~uint
	~byte
	~uint16
	~uint32
	~uint64
	~float32
	~float64
	~complex64
	~complex128
	~uintptr
	~bool
	~string
	~[3]uint8
	~[]complex128
	~map[string]uintptr
	~func(i int) (b bool)
	~struct{ a, b int }
	~chan int
}

type I30 interface {
	int
	string
}

type I31 interface {
	~int
	~string
}

type I32 interface {
	int
	string
	~int
	~string
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Number interface {
	Integer
	String() string
}

/**
命名或匿名的类型约束接口只能用于类型参数声明

不能用于全局/局部变量、函数/方法形参以及结构体字段的声明
*/

// 为函数声明类型参数T, 其类型为any
func ZeroValue[T any]() {
	var x T
	fmt.Printf("zero value of %T type: %v\n", x, x)
}

func Fn0[T, U any](t T, u U) {
	fmt.Println(t, u)
}

func Fn1[T Integer]() (t T) {
	return T(rand.Int()) // 使用类型参数强制转换类型
}

func Fn2[K comparable, V any](m map[K]V) {
	fmt.Println(m)
}

func main() {
	ZeroValue[bool]()           // zero value of bool type: false
	ZeroValue[complex64]()      // zero value of complex64 type: (0+0i)
	ZeroValue[[3]int]()         // zero value of [3]int type: [0 0 0]
	ZeroValue[map[string]int]() // zero value of map[string]int type: map[]
	ZeroValue[struct {
		b bool
		e error
	}]() // zero value of struct { b bool; e error } type: {false <nil>}

	Fn0[int, int](123, 456)                // 123 456
	Fn0[string, string]("hello", "world")  // hello
	fmt.Println(Fn1[int64]())              // 5577006791947779410
	Fn2[int, int](map[int]int{1: 2, 3: 4}) // map[1:2 3:4]

	// 函数调用省略类型参数实参[...], Go自动推导其类型
	Fn0(789, "xyz")                             // 789 xyz
	Fn2(map[string]bool{"x": true, "y": false}) // map[x:true y:false]

	// 无法推导类型的场景不能省略类型参数实参
	fmt.Println(Fn1[uint64]()) // 8674665223082153551
}
