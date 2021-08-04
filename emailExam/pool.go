package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"sync"
	"time"
)

func main() {
	ch := make(chan *email.Email, 10) //有缓冲
	p, err := email.NewPool(
		"smtp.126.com:25",
		4, //4个链接实例
		smtp.PlainAuth("jibenliu", "jibenliu@126.com", "KFLEVYAVKWKBISLZ", "smtp.126.com"),
	)
	if err != nil {
		log.Fatal("failed to create pool:", err)
	}
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ { //4个链接池去消费
		go func() {
			defer wg.Done()
			for e := range ch {
				err := p.Send(e, 10*time.Second)
				if err != nil {
					fmt.Fprintf(os.Stderr, "email:%v sent error:%v\n", e, err)
				}
			}
		}()
	}
	for i := 0; i < 10; i++ { //创建10个email对象channel
		e := email.NewEmail()
		e.From = "hh <jibenliu@126.com>"
		e.To = []string{"jibenleo@163.com"}
		e.Subject = "Awesome web"
		e.Text = []byte(fmt.Sprintf("Awesome Web %d", i+1))
		ch <- e
	}

	close(ch)
	wg.Wait()
}
