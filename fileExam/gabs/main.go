package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{
	"info": {
		  "name": {
			"first": "lee",
			"last": "darjun"
		  },
		  "age": 18,
		  "hobbies": [
			"game",
			"programming"
		  ]
		}
	}`))

	fmt.Println("first name: ", jObj.Search("info", "name", "first").Data().(string))
	gObj, _ := jObj.JSONPointer("/info/age")
	fmt.Println("age: ", gObj.Data().(float64))
	fmt.Println("one hobby: ", jObj.Path("info.hobbies.1").Data().(string))
}


// gabs提供 3 种查询方式：读取json
//
// 以.分隔的路径调用Path()方法；
// 将路径各个部分作为可变参数传入Search()方法；
// 使用/分隔的路径调用JSONPointer()方法。