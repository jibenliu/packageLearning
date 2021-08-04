package main

import "net/http"

type Config struct {
	address string
	port    string
}

type Server struct {
	config *Config
}

func NewServer() *Server {
	return &Server{BuildConfig()}
}

func BuildConfig() *Config {
	return &Config{
		address: "127.0.0.1",
		port:    "8080",
	}
}

func main() {
	svc := NewServer()
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		_, _ = resp.Write([]byte("di"))
	})
	_ = http.ListenAndServe(svc.config.address+":"+svc.config.port, nil)
}
