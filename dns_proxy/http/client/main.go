package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Getenv("HTTP_PROXY"))
	u, err := url.Parse(os.Getenv("HTTP_PROXY"))
	if err != nil {
		fmt.Println(err, u)
	}
	resolver := &net.Resolver{
		PreferGo: true, //否则不生效
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.DialUDP("udp", nil, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 53})
		},
	}

	fmt.Println(resolver.LookupAddr(context.TODO(), net.ParseIP("127.0.0.8").String()))
	//[xiazemin.com.] <nil>
	fmt.Println(resolver.LookupAddr(context.TODO(), net.ParseIP("127.0.0.1").String()))
	//[localhost kubernetes.docker.internal.] <nil>

	dialer := &net.Dialer{
		Timeout:  1 * time.Second,
		Resolver: resolver,
	}
	client := http.Client{
		Transport: &http.Transport{
			DialContext:         dialer.DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
			//	Proxy:               http.ProxyURL(u),
			Proxy: http.ProxyFromEnvironment,
		},
		Timeout: 1 * time.Second,
	}
	req, _ := http.NewRequest("GET", "http://xiazemin.com:8080", &net.Buffers{[]byte("xiazemin http get")})
	fmt.Println(client.Do(req))
}
