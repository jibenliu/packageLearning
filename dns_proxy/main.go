package main

import (
	"sync"

	"github.com/xiazemin/dns_proxy/dns"
	"github.com/xiazemin/dns_proxy/http/server"
	"github.com/xiazemin/dns_proxy/tcpproxy"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		dns.Serve()
	}()
	go func() {
		defer wg.Done()
		tcpproxy.Serve()
	}()
	go func() {
		defer wg.Done()
		server.ExampleServe()
	}()
	wg.Wait()
}
