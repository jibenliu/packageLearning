/**
 * benchMark测试
 * go test string_benchmark.go -benchmem 输出结果
 */
package main

import (
	"bytes"
	"strings"
	"testing"
)

const (
	sss = "https://mojotv.cn"
	cnt = 10000
)

var (
	bbb      = []byte(sss)
	expected = strings.Repeat(sss, cnt)
)

//使用 提前初始化  内置 copy函数
func BenchmarkCopyPreAllocate(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		bs := make([]byte, cnt*len(sss))
		bl := 0
		for i := 0; i < cnt; i++ {
			bl += copy(bs[bl:], sss)
		}
		result = string(bs)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result;got=%s,want=%s", string(result), expected)
	}
}

//使用 提前初始化  内置append 函数
func BenchmarkAppendPreAllocate(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		data := make([]byte, 0, cnt*len(sss))
		for i := 0; i < cnt; i++ {
			data = append(data, sss...)
		}
		result = string(data)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 提前初始化 bytes.Buffer
func BenchmarkBufferPreAllocate(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		buf := bytes.NewBuffer(make([]byte, 0, cnt*len(sss)))
		for i := 0; i < cnt; i++ {
			buf.WriteString(sss)
		}
		result = buf.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 strings.Repeat 本质是pre allocate + strings.Builder
func BenchmarkStringRepeat(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = strings.Repeat(sss, cnt)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 内置copy
func BenchmarkCopy(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		data := make([]byte, 0, 64) // same size as bootstrap array of bytes.Buffer
		for i := 0; i < cnt; i++ {
			off := len(data)
			if off+len(sss) > cap(data) {
				temp := make([]byte, 2*cap(data)+len(sss))
				copy(temp, data)
				data = temp
			}
			data = data[0 : off+len(sss)]
			copy(data[off:], sss)
		}
		result = string(data)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 内置append
func BenchmarkAppend(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		data := make([]byte, 0, 64)
		for i := 0; i < cnt; i++ {
			data = append(data, sss...)
		}
		result = string(data)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 bytes.Buffer
func BenchmarkBufferWriteBytes(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		for i := 0; i < cnt; i++ {
			buf.Write(bbb)
		}
		result = buf.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

func BenchmarkStringBuilderWriteBytes(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		var buf strings.Builder
		for i := 0; i < cnt; i++ {
			buf.Write(bbb)
		}
		result = buf.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}
func BenchmarkBufferWriteString(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		for i := 0; i < cnt; i++ {
			buf.WriteString(sss)
		}
		result = buf.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

func BenchmarkStringPlusOperator(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		var str string
		for i := 0; i < cnt; i++ {
			str += sss
		}
		result = str
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}
