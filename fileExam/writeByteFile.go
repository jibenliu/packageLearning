package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type Article struct {
	Id      int
	Title   string
	Content string
	Author  string
}

// 写入二进制数据到磁盘文件
func write(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// 从磁盘文件加载二进制数据
func read(data interface{}, filename string) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	article := Article{
		Id:      1,
		Title:   "基于 Gob 包编解码二进制数据",
		Content: "通过 Gob 包序列化二进制数据以便通过网络传输",
		Author:  "学院君",
	}
	write(article, "article_data")
	var articleData Article
	read(&articleData, "article_data")
	fmt.Printf("%#v\n", articleData)
}
