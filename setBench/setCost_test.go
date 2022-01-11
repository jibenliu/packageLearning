package setBench

import (
	"testing"
)

const num = int(1 << 24)

func BenchmarkSetWithBoolValueWrite(b *testing.B) {
	set := make(map[int]bool)
	for i := 0; i < num; i++ {
		set[i] = true
	}
}

// BenchmarkSetWithInterfaceValueWrite 测试 interface{} 类型
func BenchmarkSetWithInterfaceValueWrite(b *testing.B) {
	set := make(map[int]interface{})
	for i := 0; i < num; i++ {
		set[i] = struct{}{}
	}
}

// BenchmarkSetWithIntValueWrite 测试 int 类型
func BenchmarkSetWithIntValueWrite(b *testing.B) {
	set := make(map[int]int)
	for i := 0; i < num; i++ {
		set[i] = 0
	}
}

// BenchmarkSetWithStructValueWrite 测试 struct{} 类型
func BenchmarkSetWithStructValueWrite(b *testing.B) {
	set := make(map[int]struct{})
	for i := 0; i < num; i++ {
		set[i] = struct{}{}
	}
}
