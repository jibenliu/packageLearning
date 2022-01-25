package main

import (
	"log"
	"net"
	"time"
)

// HeartBeating , determine if client send a message within set time by GravelChannel
// 心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息
func HeartBeating(conn net.Conn, readerChannel chan byte, timeout int) {
	select {
	case _ = <-readerChannel:
		log.Printf(conn.RemoteAddr().String(), "get message, keeping heartbeating...")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(time.Second * 5):
		log.Printf("It's really weird to get Nothing!!!")
		conn.Close()
	}
}
