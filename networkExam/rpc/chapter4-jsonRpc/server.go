package main

import (
	"encoding/json"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	mp := make(map[string]string)
	mp["hello"] = request
	replyStr, _ := json.Marshal(mp)
	*reply = string(replyStr)
	return nil
}

func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}


/**
	启动前
		nc -l 1234 然后调用client 查看客户端调用时发送的数据格式
·	启动后
		模拟client请求 echo -e '{"method":"HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234
 */