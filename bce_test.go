package main

import (
	"strings"
	"testing"
)

func NumSameBytes_1(x, y string) int {
	if len(x) > len(y) {
		x, y = y, x
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return i
		}
	}
	return len(x)
}

func NumSameBytes_2(x, y string) int {
	if len(x) > len(y) {
		x, y = y, x
	}
	if len(x) <= len(y) { // 虽然代码多了，但是效率提高了
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] { // 边界检查被消除了
				return i
			}
		}
	}
	return len(x)
}

var x = strings.Repeat("hello", 100) + " world!"
var y = strings.Repeat("hello", 99) + " world!"

func BenchmarkNumSameBytes_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumSameBytes_1(x, y)
	}
}

func BenchmarkNumSameBytes_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumSameBytes_2(x, y)
	}
}

// 在上级目录下  go test --bench=。 bce/bce_test.go -benchmem