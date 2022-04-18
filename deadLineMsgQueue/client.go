// 消费者.go
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@127.0.0.1:8161/")
	if err != nil {
		log.Fatal(err)
	}

	c, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := c.Consume("dlxQueue", "", true, false, false, false, nil) //监听dlxQueue队列
	if err != nil {
		log.Fatal(err)
	}

	for d := range msgs {
		fmt.Printf("收到信息: %s\n", d.Body) // 收到消息，业务处理
	}
}

// 5秒之后，打印
// 收到信息: hello world1
// 收到信息: hello world1
// 收到信息: hello world1
// 收到信息: hello world1
// 收到信息: hello world1
// 收到信息: hello world1
// 收到信息: hello world1
// 收到信息: hello world1
// 收到信息: hello world1
// 收到信息: hello world1
