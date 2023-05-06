package main

import (
	"bufio"
	"fmt"
	"github.com/gamexg/proxyclient"
)

func main() {
	p, err := proxyclient.NewProxyClient("socks5://root:password@111.67.195.179:9898")
	if err != nil {
		panic(err)
	}

	c, err := p.Dial("tcp", "www.accebits.com:80")
	if err != nil {
		panic(err)
	}

	// 发送 HTTP 请求
	req := "GET / HTTP/1.1\r\n" +
		"Host: accebits.com\r\n" +
		"Connection: close\r\n\r\n"
	fmt.Fprintf(c, req)

	// 读取响应
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
