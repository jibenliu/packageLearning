## pill.go文件是手写的，pill_string.go文件是执行go generate 命令生成的

#### 为的是实现增加常量时不改动代码，实现string方法


##### make命令实现执行go generate同时编译项目了
```bigquery
all:
    go generate && go build .
```