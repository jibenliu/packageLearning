package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@127.0.0.1:8161/")
	if err != nil {
		log.Fatal("dial error: ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	args := amqp.Table{"x-dead-letter-exchange": "dlx"}
	q, err := ch.QueueDeclare("test", true, false, false, false, args) // 声明一个test队列，并设置队列的死信交换机为"dlx"

	body := "hello world1"
	for i := 0; i < 10; i++ {
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			Body:       []byte(body),
			Expiration: "5000", // 设置TTL为5秒
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
