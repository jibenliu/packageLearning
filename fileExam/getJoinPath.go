package main

import "path/filepath"

func main() {
	//拼装路径
	path := filepath.Join(`./a`, `b/c`, `../d/`, `e`)
	println(path)
}
