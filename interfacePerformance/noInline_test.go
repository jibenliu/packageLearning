package interfacePerformance

import "testing"

//go:noinline
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Result int

func BenchmarkMax(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = max(-1, i)
	}
	Result = r
}

// 内联是一个基本的编译器优化，它用被调用函数的主体替换函数调用。以消除调用开销，但更重要的是启用了其他编译器优化
//		提升主要有两方面
//			消除了函数调用本身的开销
//			允许编译器更有效地应用其他优化策略（例如常量折叠，公共子表达式消除，循环不变代码移动和更好的寄存器分配）
// go test -bench=. -benchmem -run=none
// 允许 max 函数内联，也就是把 //go:noinline 这行代码删除  重新测试

