package main

import (
	"fmt"
	"net"
)

var nick = ""

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err == nil {
		_ = fmt.Sprintf("connection failed: %+v", err)
	}

	defer conn.Close()
	fmt.Println("connect server succeed!")

	fmt.Printf("Make a nickname for your chatRoom: ")
	_, _ = fmt.Scanf("%s", &nick)
	fmt.Println("hello : ", nick)

	_, _ = conn.Write([]byte("nick|" + nick))

	go clientHandle(conn)

	var msg string

	for {
		msg = ""
		_, _ = fmt.Scan(&msg)

		_, _ = conn.Write([]byte("say|" + nick + "|" + msg))

		if msg == "quit" {
			_, _ = conn.Write([]byte("quit|" + nick))
			break
		}
	}
}

func clientHandle(conn net.Conn) {
	for {
		data := make([]byte, 255)
		msgRead, err := conn.Read(data)
		if msgRead == 0 || err != nil {
			break
		}
		fmt.Println(string(data[0:msgRead]))
	}
}
