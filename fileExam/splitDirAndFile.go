package main

import "path/filepath"

func main() {
	dir, file := filepath.Split("/Users/jibenliu/data/dnmp/serverStart.bat")
	println(dir)
	println(file)
}
