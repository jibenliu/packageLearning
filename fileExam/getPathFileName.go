package main

import "path/filepath"

func main() {
	//获取路径的文件名
	file := filepath.Base(`a/b/c/d.go`)
	println(file)

	path := filepath.Base(`a/b/c`)
	println(path)
}
