package main

import (
	"os"
	"path/filepath"
)

func main() {
	//在一些 特殊使用场景，上述的方法 会获取到错误路径
	//目前已知：系统来调用这个可执行文件
	dir, _ := os.Getwd()
	println(dir)

	//替代方法
	dirs, _ := os.Executable()
	exPath := filepath.Dir(dirs)
	println(exPath)

}
