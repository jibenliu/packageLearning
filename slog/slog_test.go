package slog

import (
	"fmt"
	"os"
	"testing"
)

func Test_log(T *testing.T) {
	File, _ := os.Create("log")
	log, err := NewLog("Info", true, File, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer log.Close()
	for i := 0; i < 10; i++ {
		log.Warn("Nima")
		log.Info("Fuck")
	}

}
func Benchmark_log(b *testing.B) {
	File, _ := os.Create("log")
	log, err := NewLog("Info", true, File, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer log.Close()
	for i := 0; i < b.N; i++ {
		log.Warn("Nima")
	}
}

// go test

//go test -bench=.
/** B类型也有以下参数：
		benchmem：输出内存分配统计
		benchtime：指定测试时间
		cpu：指定GOMAXPROCS
		timeout：超时限制
			go test -v -bench=. -cpu=8 -benchtime="3s" -timeout="5s" -benchmem
				输出：
				goos: darwin
				goarch: amd64
				Benchmark-8     5000000000           0.34 ns/op        0 B/op          0 allocs/op
				PASS
				ok      _/Users/golang_learning/testTB  1.766s
			Benchmark-8：-cpu参数指定，-8表示8个CPU线程执行
				5000000000：表示总共执行了5000000000次
				0.34 ns/op：表示每次执行耗时0.34纳秒
				0 B/op:表示每次执行分配的内存（字节）
				0 allocs/op：表示每次执行分配了多少次对象

	命令行生成测试数据文件
		$ go test -bench=. -cpuprofile cpu.out
		输出：
		goos: darwin
		goarch: amd64
		Benchmark-4     2000000000           0.35 ns/op
		PASS
		ok      _/Users/golang_learning/testTB  0.911s
		$ ls
		输出：（上一条命令生成cpu.out和testTB.test）
		add_test.go     cpu.out     testTB.test
	go tool pprof -text mem.out
		输出：
		Main binary filename not available.
		Type: inuse_space
		Time: May 22, 2018 at 3:36pm (CST)
		Showing nodes accounting for 1.16MB, 100% of 1.16MB total
			  flat  flat%   sum%        cum   cum%
			1.16MB   100%   100%     1.16MB   100%  runtime/pprof.StartCPUProfile
				 0     0%   100%     1.16MB   100%  main.main
				 0     0%   100%     1.16MB   100%  runtime.main


*/
