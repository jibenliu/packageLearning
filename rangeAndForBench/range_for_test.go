package rangeAndForBench

import "testing"

func BenchmarkRangeHiPerformance(b *testing.B) {
	v := CreateABigSlice(1 << 12)

	for i := 0; i < b.N; i++ {
		length := len(v)
		var tmp [4096]int
		for k := 0; k < length; k++ {
			tmp = v[k]
		}
		_ = tmp
	}
}

func BenchmarkRangeLowPerformance(b *testing.B) {
	v := CreateABigSlice(1 << 12)

	for i := 0; i < b.N; i++ {
		var tmp [4096]int
		for _, e := range v {
			tmp = e
		}
		_ = tmp
	}
}

// go test -bench=.

// 经检查 range 遇到元素占用的内存比较大时，和 for 性能接近，测试版本1.19.3
