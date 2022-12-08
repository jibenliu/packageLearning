#refer:[https://darjun.github.io/2022/11/15/godailylib/ebiten1/](https://darjun.github.io/2022/11/15/godailylib/ebiten1/)

# 将图片打包进代码
> go generate
----

# 1.本地运行 
> go run .

----


# 2.wasm运行
## 2.1先编译
> GOOS=js GOARCH=wasm go build -o alien_invasion.wasm
## 2.2 添加server服务器
将位于$GOROOT/misc/wasm目录下的wasm_exec.html和wasm_exec.js文件拷贝到我们的项目目录下。注意wasm_exec.html文件中默认是加载名为test.wasm的文件，我们需要将加载文件改为alien_invasion.wasm，或者将生成的文件改名为test.wasm
```golang
package main

import (
  "log"
  "net/http"
)

func main() {
  if err := http.ListenAndServe(":8080", http.FileServer(http.Dir("."))); err != nil {
    log.Fatal(err)
  }
}
```
## 2.3 访问
> 打开浏览器输入地址：localhost:8080/wasm_exec.html