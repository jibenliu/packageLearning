#refer:[https://www.syst.top/posts/go/handle-million-requests/](https://www.syst.top/posts/go/handle-million-requests/)


```shell
安装 github.com/link1st/go-stress-testing
```

# DOC
- -c 表示并发数

- -n 每个并发执行请求的次数，总请求的次数 = 并发数 * 每个并发执行请求的次数

- -u 需要压测的地址
```shell
 go-stress-testing-linux -c 1000 -n 50000 -u http://localhost:8099/payload
```
