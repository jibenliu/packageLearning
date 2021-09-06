##查看内存分配
> go test -bench=.

##查看各自耗时
> go test -bench=Or

>go test -bench=InReflect 


### 生成测试的应用程序
```golang
go test -c
mv orDone.test bench_test
```

### 运行测试文件
```
./bench_test -test.bench=Or 
```