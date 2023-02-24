refer:[实战 Go：如何实现一个简单分布式系统](https://juejin.cn/post/7158685917006266398?share_token=63630477-4ab5-4cdb-97af-6c68c2b63315#heading-1)

```shell
go run main.go master

go run main.go


curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"cmd": "touch ./tmp/hello-distributed-system"}' \
  http://localhost:9092/tasks
```