package main

import (
	"context"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/ybbus/httpretry"
)

func main() {
	client := httpretry.NewDefaultClient(
		httpretry.WithMaxRetryCount(10),
		httpretry.WithRetryPolicy(func(statusCode int, err error) bool {
			return err != nil || statusCode >= 500 || statusCode == 0
		}),
		httpretry.WithBackoffPolicy(func(attemptNum int) time.Duration { //指数退避策略
			return time.Duration(math.Pow(2, float64(attemptNum)))*time.Second + time.Duration(rand.Int63n(int64(time.Second)))/2
		}),
	)
	url := "https://api.example.com/v1/thing/123"
	var errRes interface{}
	var resp interface{}
	err := requests.URL(url).Client(client).Method(http.MethodPut).BodyJSON(map[string]interface{}{
		"fields": map[string]interface{}{
			"abc": "def",
		},
	}).Bearer("token").ErrorJSON(&errRes).ToJSON(&resp).Fetch(context.Background())
	if err != nil || errRes != nil {
		log.Printf("send request error:%v, errRes:%v, resp:%v", err, errRes, resp)
	}
}
