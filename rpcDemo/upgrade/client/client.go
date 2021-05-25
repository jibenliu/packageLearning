package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	defer client.Close()

	var result Result
	asyncCall := client.Go("Cal.Square", 12, &result, nil)
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	<-asyncCall.Done
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
