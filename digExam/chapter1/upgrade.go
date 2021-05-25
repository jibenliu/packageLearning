package main

import "net/http"

type Config struct {
	address string
	port    string
}

type Server struct {
	config *Config
}

func NewServer(c *Config) *Server {
	return &Server{c}
}

func main() {
	c := Config{"127.0.0.1", "8080"}
	svc := NewServer(&c)
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		_, _ = resp.Write([]byte("di"))
	})
	_ = http.ListenAndServe(svc.config.address+":"+svc.config.port, nil)
}