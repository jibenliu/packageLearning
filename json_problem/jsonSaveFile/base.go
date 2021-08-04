package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  string `json:"author"`
}

func main() {
	var books map[int]*Book = make(map[int]*Book)
	book1 := Book{Id: 1, Title: "Go web编程", Summary: "Go web编程入门指南", Author: "学院君"}
	books[book1.Id] = &book1

	// 相当于file_put_contents
	//// 通过 JSON 序列化字典数据
	//data, _ := json.Marshal(books)
	//
	//// 将 JSON 格式数据写入当前目录下的d books 文件（文件不存在会自动创建）
	//err := ioutil.WriteFile("books", data, 0644)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 从文件 books 中读取数据
	//dataEncoded, _ := ioutil.ReadFile("books")
	//// 将读取到的数据通过 JSON 解码反序列化为原来的数据类型
	//var booksDecoded map[int]*Book
	//json.Unmarshal(dataEncoded, &booksDecoded)
	//fmt.Printf("%#v", booksDecoded[book1.Id])


	// 相当于fopen fwrite
	// 通过 JSON 序列化字典数据
	data, _ := json.Marshal(books)
	// err := ioutil.WriteFile("books", data, 0644)
	file1, _ := os.Create("books")
	defer file1.Close()
	_, err := file1.Write(data)
	if err != nil {
		panic(err)
	}

	// 从文件 books 中读取数据
	//dataEncoded, _ = ioutil.ReadFile("books")
	file2, _ := os.Open("books")
	defer file2.Close()
	dataEncoded := make([]byte, len(data))
	file2.Read(dataEncoded)
	// 将读取到的数据通过 JSON 解码反序列化为原来的数据类型
	var booksDecoded map[int]*Book
	json.Unmarshal(dataEncoded, &booksDecoded)
	fmt.Printf("%#v", booksDecoded[book1.Id])

}
