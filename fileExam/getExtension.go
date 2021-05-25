package main

import "path/filepath"

func main() {
	//获取文件扩展名
	fileExt := filepath.Ext(`/Users/jibenliu/data/dnmp/serverStart.bat`)

	println(fileExt)
	println(fileExt[1:])

	fileExt2 := filepath.Ext(`/Users/jibenliu/data/dnmp/README`)

	println(fileExt2)
	//println(fileExt2[1:])
}


