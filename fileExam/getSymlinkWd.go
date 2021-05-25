package main

import "path/filepath"

func main() {
	//获取链接文件的真实路径

	path, err := filepath.EvalSymlinks(`data`)
	if err != nil {
		panic(err)
	}
	println(path)
}
