package main

import "fmt"

func SumVariables(i, j int, quit chan bool) {
	z := i + j
	fmt.Println(z)
	quit <- true
}

func main() {
	channels := make([]chan bool, 10)
	for i := 0; i < 10; i++ {
		channels[i] = make(chan bool)
		go SumVariables(i, i, channels[i])
	}
	for _, v := range channels {
		<-v
	}
}
