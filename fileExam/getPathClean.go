package main

import "path/filepath"

func main() {
	//清理路径
	path := filepath.Clean(`./a/b/../c/d/e/../f`)
	println(path)
}
