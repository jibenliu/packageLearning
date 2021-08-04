package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var slice1 []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "张三"
	m1["address"] = "广东省深圳市"
	slice1 = append(slice1, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "李四"
	m2["address"] = "广东省广州市"
	slice1 = append(slice1, m2)

	data, err := json.Marshal(slice1)
	if err != nil {
		fmt.Println("序列化错误")
	} else {
		fmt.Println(string(data))
	}
}
