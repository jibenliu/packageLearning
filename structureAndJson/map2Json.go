package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var map1 map[string]interface{} // 使用一个空接口表示可以是任意类型
	map1 = make(map[string]interface{})
	map1["name"] = "张三"
	map1["province"] = "广东省"
	map1["city"] = "深圳市"
	map1["salary"] = 1000
	map1["hobby"] = [...]string{"看书", "旅游", "学习"}

	data, err := json.Marshal(map1)
	if err != nil {
		fmt.Println("序列化失败")
	} else {
		fmt.Println(string(data))
	}
}
