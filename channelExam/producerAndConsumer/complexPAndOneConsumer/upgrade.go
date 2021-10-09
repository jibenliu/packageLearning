package main

import "fmt"

type Msg struct {
	in int
}

func producer(in []int) chan Msg {
	ch := make(chan Msg)
	go func() {
		for _, v := range in {
			ch <- Msg{in: v}
		}
		close(ch)
	}()
	return ch
}

func consumer(ch1, ch2 chan Msg) {
	var v1 Msg
	var v2 Msg
	ok1 := true
	ok2 := true

	for ok1 || ok2 {
		select {
		case v1, ok1 = <-ch1:
			fmt.Println(v1)
			if !ok1 { //通道关闭了
				ch1 = nil
			}
		case v2, ok2 = <-ch2:
			fmt.Println(v2)
			if !ok2 { //通道关闭了
				ch2 = nil
			}
		}
	}
}

func mpsc() {
	ch1 := producer([]int{1, 2, 3})
	ch2 := producer([]int{4, 5, 6})

	consumer(ch1, ch2)
}
func main() {
	mpsc()
}
