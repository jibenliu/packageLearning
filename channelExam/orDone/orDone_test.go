package orDone

import "testing"

func BenchmarkOr(b *testing.B) {
	done := make(chan interface{})
	defer close(done)
	num := 100
	streams := make([]<-chan interface{}, num)
	for i := range streams {
		streams[i] = repeat(done, []int{1, 2, 3})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-Or(streams...)
	}
}

func BenchmarkOrInReflect(b *testing.B) {
	done := make(chan interface{})
	defer close(done)
	num := 100
	streams := make([]<-chan interface{}, num)
	for i := range streams {
		streams[i] = OrInReflect(done)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-Or(streams...)
	}
}
