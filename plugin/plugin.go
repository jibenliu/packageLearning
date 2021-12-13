package main

import (
	"log"
)

func init() {
	log.Println("plugin init")
}

type SayHello struct {
}

func (s *SayHello) CallMe(name string) string {
	log.Println("hello ", name)
	return "I am plugin"
}

// SayHelloPlugin 导出变量
var SayHelloPlugin SayHello

// go build -o plugin.so -buildmode=plugin plugin.go
