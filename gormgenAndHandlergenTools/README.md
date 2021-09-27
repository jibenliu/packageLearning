

## gormgen和handlergen 两个工具的安装使用


### gormgen
```go
go get -u github.com/MohamedBassem/gormgen/...
```
- Add the //go:generate comment mentioned above anywhere in your code.
- Add go generate to your build steps.
- The generated code will depend on gorm and gormgen, so make sure to vendor both of them.


### handlergen
```

go install github.com/Karitham/handlergen@latest


handlergen -file handler.yaml > hander_generated.go
```

