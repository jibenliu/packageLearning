package main

import "testing"

const N = 1024 * 1024

type Element = int64

var xForMake = make([]Element, N)
var xForMakeCopy = make([]Element, N)
var xForAppend = make([]Element, N)
var yForMake []Element
var yForMakeCopy []Element
var yForAppend []Element

func Benchmark_PureMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMake = make([]Element, N)
	}
}

func Benchmark_MakeAndCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMakeCopy = make([]Element, N)
		copy(yForMakeCopy, xForMakeCopy)
	}
}

func Benchmark_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForAppend = append([]Element(nil), xForAppend...)
	}
}

// go test -bench=. -benchtime=3s


// 在go12 使用append来克隆切片比使用make+copy要高效得多。使用make+copy相对低效的原因是make需要将其开辟出的每个元素置零，这对于这个应用场景其实是没有必要的
// 在go15 make+copy甚至比一个单make都快（因为单make需要对元素清零，而且此清零操作尚未被优化）
