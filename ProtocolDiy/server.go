package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	addr := fmt.Sprintf("%s:%s", "127.0.0.1", "8090")
	listen, err1 := net.Listen("tcp", addr)

	if err1 != nil {
		fmt.Println("Error:%s", err1.Error())
		os.Exit(1)
	}

	for {
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("Error:%s", err2.Error())
			os.Exit(1)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	msg := make(chan string, ChanMsgCount)

	buffer1 := NewBuffer(conn, HEADER, BufferLength)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		doMsg(msg)
		defer wg.Add(-1)
	}()

	go func() {
		err := buffer1.Read(msg)
		if err != nil {
			if err.Error() == "EOF" {
				close(msg)
			} else {
				panic(err)
			}
		}
		defer wg.Add(-1)
	}()

	wg.Wait()
	fmt.Println("一个客户端处理的消息处理完毕")
}

func doMsg(msg chan string) {
	count := 0
	for v := range msg {
		fmt.Println("第", count, "个消息体长：", len(v))
		count++
	}
}
