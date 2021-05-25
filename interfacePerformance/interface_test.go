package interfacePerformance

import "testing"

type D interface {
	Append(D)
}

type Strings []string

func (s Strings) Append(d D) {

}

func BenchmarkInterface(b *testing.B) {
	s := D(Strings{})
	for i := 0; i < b.N; i += 1 {
		s.Append(Strings{""})
	}
}

func BenchmarkConcrete(b *testing.B) {
	s := Strings{}
	for i := 0; i < b.N; i += 1 {
		s.Append(Strings{""})
	}
}

func BenchmarkInterfaceTypeAssert(b *testing.B) {
	s := D(Strings{})
	for i := 0; i < b.N; i += 1 {
		s.(Strings).Append(Strings{""})
	}
}


// bench测试
// go test -bench=. -benchmem


// BenchmarkInterface-4                    14717325                74.0 ns/op            48 B/op          2 allocs/op
// 显示第一条执行了14717325次，平均耗时74ns 平均使用内存48B 每次操作进行2次内存分配


// brew install graphviz
// refer: https://www.cnblogs.com/nickchen121/p/11517452.html#%E4%BA%94%E3%80%81go-tool-pprof%E5%91%BD%E4%BB%A4
// 根据pprof分析
// go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
// go tool pprof profile.out
// 		进入pprof
// 		然后你也可以用list命令检查函数需要的时间  list logicCode
// 		使用web命令生成图像（png,pdf,...） png 1.png
// 或者 go tool pprof -pdf profile_cpu.out > profile_cpu.pdf