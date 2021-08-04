package main

import (
	"log"
	"net"
	"net/rpc"
)

const sHelloServiceName = "path/to/pkg.HelloService"

type HelloService struct{}
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(sHelloServiceName, svc)
}

func main() {
	_ = RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
