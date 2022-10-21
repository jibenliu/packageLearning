package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://root:123456@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 获取channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable 持久化，当为true时可以让mq在重启后消息不丢失 ，持久的选项更改需要同时应用于生产者代码和消费者代码
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 告诉RabbitMQ不要一次向一个worker发出多个消息。或者，换句话说，在处理并确认前一条消息之前，不要向worker发送新消息。相反，它将把它发送给下一个不忙的worker
	// 防止因配置或者网络问题导致消费能力不同而造成消息均衡被发布阻塞
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		fmt.Printf("ch.Qos() failed, err:%v\n", err)
		return
	}

	// 获取接收消息的Delivery通道
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack 需要主动确认才算消费成功
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte(".")) // 数一下有几个.
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second) // 模拟耗时的任务
			log.Printf("Done")
			d.Ack(false) // 手动传递消息确认
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

/**
go run main.go msg1. &&
go run main.go msg2.. &&
go run main.go msg3... &&
go run main.go msg4.... &&
go run main.go msg5..... &&
go run main.go msg6......
*/
