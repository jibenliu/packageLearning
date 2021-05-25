package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

//不对服务端鉴权
//func main() {
//	config := &tls.Config{
//		InsecureSkipVerify: true,
//	}
//	conn, _ := tls.Dial("tcp", "localhost:1234", config)
//	defer conn.Close()
//	client := rpc.NewClient(conn)
//
//	var result Result
//	if err := client.Call("Cal.Square", 12, &result); err != nil {
//		log.Fatal("Failed to call Cal.Square. ", err)
//	}
//
//	log.Printf("%d^2 = %d", result.Num, result.Ans)
//}

//开启对服务端鉴权
func main()  {
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile("../server/server.crt")
	if err != nil {
		log.Fatal("Failed to read server.crt")
	}
	certPool.AppendCertsFromPEM(certBytes)

	//服务端如果需要对客户端鉴权
	//cert, _ := tls.LoadX509KeyPair("client.crt", "client.key")
	//certPool := x509.NewCertPool()
	//certBytes, _ := ioutil.ReadFile("../server/server.crt")
	//certPool.AppendCertsFromPEM(certBytes)
	//config := &tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//	RootCAs: certPool,
	//}

	config := &tls.Config{
		RootCAs: certPool,
	}

	conn, _ := tls.Dial("tcp", "localhost:1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)

	var result Result
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}

	log.Printf("%d^2 = %d", result.Num, result.Ans)
}