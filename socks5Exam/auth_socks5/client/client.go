package main

import (
	"fmt"
	"github.com/gamexg/proxyclient"
	"io"
)

func main() {
	p, err := proxyclient.NewProxyClient("socks5://root:password@127.0.0.1:8888")
	if err != nil {
		panic(err)
	}

	c, err := p.Dial("tcp", "www.accebits.com:443")
	if err != nil {
		panic(err)
	}

	io.WriteString(c, "GET /index.html HTTP/1.0\r\nHOST:www.163.com\r\n\r\n")
	b, err := io.ReadAll(c)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
