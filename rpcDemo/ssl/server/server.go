package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	//服务端需要对客户端鉴权
	//cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	//certPool := x509.NewCertPool()
	//certBytes, _ := ioutil.ReadFile("../client/client.crt")
	//certPool.AppendCertsFromPEM(certBytes)
	//config := &tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//	ClientAuth:   tls.RequireAndVerifyClientCert,
	//	ClientCAs:    certPool,
	//}

	rpc.Register(new(Cal))
	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", ":1234", config)
	log.Printf("Serving RPC server on port %d", 1234)

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
