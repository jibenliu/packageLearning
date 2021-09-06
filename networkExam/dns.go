package main

import (
	"fmt"
	"net"
)

func main() {
	nameserver, _ := net.LookupNS("baidu.com")
	for _, ns := range nameserver {
		fmt.Println(ns)
	}
}
