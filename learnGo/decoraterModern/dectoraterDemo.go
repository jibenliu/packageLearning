package main

import (
	"fmt"
	"log"
	"time"
)

func myProcess() {
	fmt.Println("hello world")
	time.Sleep(1 * time.Second)
}

func myLog(a func()) {
	log.Printf("Starting function execution:%s\n", time.Now())
	a()
	log.Printf("Ending function execution:%s\n", time.Now())
}

func main() {
	fmt.Printf("Type:%T\n", myLog)
	myLog(myProcess)
}
