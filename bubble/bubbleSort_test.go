package bubble

import "testing"

func TestBubbleSort(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("BubbleSort() failed.Got ", values, "Expected 1 2 3 4 5")
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 ||
		values[4] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 5 5")
	}
}

func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	BubbleSort(values)
	if values[0] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 5")
	}
}

// go test --bench=sort my_test.go -benchmem
// -count n 跑n次benchmark，n默认为1
// -benchmem 打印内存分配的信息
// -benchtime=5s 自定义测试时间，默认为1s
func BenchmarkSort(b *testing.B) {
	values := []int{5, 4, 3, 2, 1}
	for n := 0; n < b.N; n++ {
		BubbleSort(values)
	}
}
