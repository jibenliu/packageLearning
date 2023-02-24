package main

import (
	"fmt"
	"os"
)

// 目录名称将会以pattern开头，一串随机数字结尾。如果pattern中包含星号*，则那串随机数字会放在星号的位置，而不是放在结尾
func main() {
	name, err := os.MkdirTemp(`./tempDir`, "temp*Dir") //创建随机名称目录
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(name)                                        // 输出：D:\temp731435683Dir
	file, err := os.CreateTemp(`./tempDir`, "temp*File.tmp") //创建随机名称文件
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file.Name())
}
