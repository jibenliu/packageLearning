package main

import (
	"fmt"
	"net"
	"strings"
)

var ConnMap = make(map[string]net.Conn)

func main() {
	listenSocket, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		_ = fmt.Sprintf("Server error:%+v", err)
	}

	defer listenSocket.Close()
	fmt.Println("Server is waiting...")

	for {
		conn, err := listenSocket.Accept()
		if err != nil {
			_ = fmt.Sprintf("conn failed:%+v", err)
		}
		fmt.Println(conn.RemoteAddr(), "connect client success!")

		go serverHandle(conn)
	}
}

func serverHandle(conn net.Conn) {
	for {
		data := make([]byte, 255)
		msgRead, err := conn.Read(data)
		if msgRead == 0 || err != nil {
			continue
		}

		msgStr := strings.Split(string(data[0:msgRead]), "|")
		switch msgStr[0] {
		case "nick":
			fmt.Println(conn.RemoteAddr(), "-->", msgStr[1])
			for k, v := range ConnMap {
				if k != msgStr[1] {
					_, _ = v.Write([]byte("[" + msgStr[1] + "] : join..."))
				}
			}
			ConnMap[msgStr[1]] = conn
		case "say":
			for k, v := range ConnMap {
				if k != msgStr[1] {
					fmt.Println("Send "+msgStr[2]+" to ", k)
					_, _ = v.Write([]byte("[" + msgStr[1] + "]: " + msgStr[2]))
				}
			}
		case "quit":
			for k, v := range ConnMap {
				if k != msgStr[1] {
					_, _ = v.Write([]byte("[" + msgStr[1] + "]: quit"))
				}
			}
			delete(ConnMap, msgStr[1])
		}
	}
}
