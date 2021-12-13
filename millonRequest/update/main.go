package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const MaxQueue = 400

var Queue chan Payload

func init() {
	Queue = make(chan Payload, MaxQueue)
}

type Payload struct {
	// 传啥不重要
}

func (p *Payload) UpdateToS3() error {
	//存储逻辑,模拟操作耗时
	time.Sleep(500 * time.Millisecond)
	fmt.Println("上传成功")
	return nil
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	// 业务过滤
	// 请求body解析......
	var p Payload
	//go p.UpdateToS3()
	Queue <- p
	w.Write([]byte("操作成功"))
}

// 处理任务
func StartProcessor() {
	for {
		select {
		case payload := <-Queue:
			payload.UpdateToS3()
		}
	}
}

func main() {
	http.HandleFunc("/payload", payloadHandler)
	//单独开一个g接收与处理任务
	go StartProcessor()
	log.Fatal(http.ListenAndServe(":8099", nil))
}
