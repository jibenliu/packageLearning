package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Tutorial struct {
	Id      int
	Title   string
	Summary string
	Author  string
}

func main() {
	csvFile, err := os.Create("tutorial.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	_, _ = csvFile.WriteString("\xEF\xBB\xBF")
	tutorials := []Tutorial{
		{Id: 1, Title: "Go 入门编程", Summary: "Go 基本语法和使用示例", Author: "学院君"},
		{Id: 2, Title: "Go Web 编程", Summary: "Go Web 编程入门指南", Author: "学院君"},
		{Id: 3, Title: "Go 并发编程", Summary: "通过并发编程提升性能", Author: "学院君"},
		{Id: 4, Title: "Go 微服务开发", Summary: "基于 go-micro 框架开发微服务", Author: "学院君"},
	}
	writer := csv.NewWriter(csvFile)
	for _, tutorial := range tutorials {
		line := []string{
			strconv.Itoa(tutorial.Id),
			tutorial.Title,
			tutorial.Summary,
			tutorial.Author,
		}
		// 将切片类型行数据写入 csv 文件
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// 将 writer 缓冲中的数据都推送到 csv 文件，至此就完成了数据写入到 csv 文件
	writer.Flush()

	// 打开这个 csv 文件
	file, err := os.Open("tutorial.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 初始化一个 csv reader，并通过这个 reader 从 csv 文件读取数据
	reader := csv.NewReader(file)
	// 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
	reader.FieldsPerRecord = -1
	// 通过 readAll 方法返回 csv 文件中的所有内容
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// 遍历从 csv 文件中读取的所有内容，并将其追加到 tutorials2 切片中
	var tutorials2 []Tutorial
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		tutorial := Tutorial{Id: int(id), Title: item[1], Summary: item[2], Author: item[3]}
		tutorials2 = append(tutorials, tutorial)
	}

	if tutorials2 != nil {
		// 打印 tutorials2 的第一个元素验证 csv 文件写入/读取是否成功
		fmt.Println(tutorials2[0].Id)
		fmt.Println(tutorials2[0].Title)
		fmt.Println(tutorials2[0].Summary)
		fmt.Println(tutorials2[0].Author)
	}
}
