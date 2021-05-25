package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var slice []map[string]interface{}
	str := `[{"address":"广东省深圳市","name":"张三"},{"address":"广东省广州市","name":"李四"}]`
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}
