package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		ticker := time.NewTicker(time.Second)
		for t := range ticker.C {
			_, err := io.WriteString(c, t.String()+"-------"+t.Format("2006-01-02 15:04:05.000")+"\n")
			if err != nil {
				return
			}
		}
	}
}

// nc localhost 8000
