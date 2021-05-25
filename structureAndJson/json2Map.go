package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var map1 map[string]interface{} // 使用一个空接口表示可以是任意类型
	//map1 = make(map[string]interface{})
	str := `{"city":"深圳市","hobby":["看书","旅游","学习"],"name":"张三","province":"广东省","salary":1000}`

	err := json.Unmarshal([]byte(str), &map1)
	if err != nil {
		fmt.Println("反序列化失败", err)
	}
	fmt.Println(map1)
}
