package main

import "fmt"

type People struct {
	name  string
	child *People
}

func main() {
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Printf("结构体：%v", relation)
}
