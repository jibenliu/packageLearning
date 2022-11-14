package tcpproxy

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
	"time"
)

//https://www.51sjk.com/b123b258404/
func Serve() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleclientrequest(client)
	}
}

func handleclientrequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	fmt.Println(method, host, address)
	hostporturl, err := url.Parse(host)
	fmt.Println(hostporturl)
	if err != nil {
		log.Println(err)
		return
	}

	if hostporturl.Opaque == "443" { //https访问
		address = hostporturl.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostporturl.Host, ":") == -1 { //host不带端口， 默认80
			address = hostporturl.Host + ":80"
		} else {
			address = hostporturl.Host
		}
	}

	//获得了请求的host和port，就开始拨号吧
	dialer := net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true, //否则不生效
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return net.DialUDP("udp", nil, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 53})
			},
		},
		Timeout: 1 * time.Second,
	}
	//net.Dial
	server, err := dialer.Dial("tcp", address)
	if err != nil {
		log.Println(err, "retry default")
		server, err = net.Dial("tcp", address)
		if err != nil {
			log.Println(err)
			return
		}
	}
	if method == "connect" {
		fmt.Fprint(client, "http/1.1 200 connection established\r\n")
	} else {
		server.Write(b[:n])
	}
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}
