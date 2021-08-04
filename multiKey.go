package main

import "fmt"

type Profile struct {
	Name    string
	Age     int
	Married bool
}

type queryKey struct {
	Name string
	Age  int
}

var mapper = make(map[interface{}]*Profile)

func buildIndex(list []*Profile) {
	for _, profile := range list {
		key := queryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}
		mapper[key] = profile
	}
}

func queryData(name string, age int) {
	key := queryKey{name, age}
	result, ok := mapper[key]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("not found")
	}
}

func main() {
	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 10},
		{Name: "王武", Age: 20},
	}

	buildIndex(list)

	queryData("张三", 30)
}
