package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func sayHello(rw http.ResponseWriter, r *http.Request) {}

func main() {
	for i := 0; i < 1000000; i++ {
		go func() {
			time.Sleep(time.Second * 10)
		}()
	}
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
