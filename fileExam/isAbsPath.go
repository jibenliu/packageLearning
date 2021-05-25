package main

import "path/filepath"

func main() {
	//判断路径是不是绝对路径
	isPath := filepath.IsAbs(`/Users/jibenliu/data/dnmp`)
	println(isPath)

	isPath2 := filepath.IsAbs(`./Users/jibenliu/data/dnmp`)
	println(isPath2)
}
