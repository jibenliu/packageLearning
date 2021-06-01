package main

import "fmt"

//go:generate echo GoGoGo!
//go:generate go run demo1.go
//go:generate echo $GOARCH $GOOS $GOFILE $GOLINE $GOPACKAGE

func main() {
	fmt.Println("go rum main.go!")
}
