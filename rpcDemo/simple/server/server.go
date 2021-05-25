package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	rpc.Register(new(Cal))
	rpc.HandleHTTP()

	log.Printf("Serving RPC server on port %d", 1234)
	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("Error serving :", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection")
}
