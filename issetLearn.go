package main

import (
	"fmt"
)

type MapContainer struct {
	items map[interface{}]interface{}
}

func (m *MapContainer) isset(index interface{}) bool {
	var target interface{}
	for k, v := range m.items {
		if k == index {
			target = v
		}
	}
	if target != nil {
		return true
	}
	return false
}
func main() {
	item := make(map[interface{}]interface{})
	item["name"] = "张三"
	item["age"] = "20"
	item["distinct"] = "武汉"
	item["area"] = "100m2"
	m := MapContainer{items: item}
	fmt.Println(m.isset("name"))
	fmt.Println(m.isset("height"))
}
