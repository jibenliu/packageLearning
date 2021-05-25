package main

import "net/http"

type Config1 struct {
	address string
	port    string
}

type Server1 struct {
	config *Config1
}

func NewServer1(c *Config1) *Server1 {
	return &Server1{c}
}

func main() {
	c := Config1{"127.0.0.1", "8080"}
	svc := NewServer1(&c)
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		_, _ = resp.Write([]byte("di"))
	})
	_ = http.ListenAndServe(svc.config.address+":"+svc.config.port, nil)
}