package main

import (
	"context"
	"fmt"
)

//go:generate echo GoGoGo!
//go:generate go run demo1.go
//go:generate echo $GOARCH $GOOS $GOFILE $GOLINE $GOPACKAGE

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("go rum main.go!")
}
