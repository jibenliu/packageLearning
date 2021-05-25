package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string   `json:"name"`   // 姓名
	Age    int      `json:"age"`    // 年龄
	Gender string   `json:"gender"` // 性别
	Score  float64  `json:"score"`  // 分数
	Course []string `json:"course"` // 课程
}

func main() {
	var stu Student
	str := `{"name":"张三","age":20,"gender":"男","score":78.6,"course":["语文","数学","音乐"]}`

	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println(stu)
}
