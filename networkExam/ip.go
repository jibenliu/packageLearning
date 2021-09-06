package main

import (
	"fmt"
	"net"
)

func main() {
	ipRecords, _ := net.LookupIP("baidu.com")
	for _, ip := range ipRecords {
		fmt.Println(ip)
	}
}
