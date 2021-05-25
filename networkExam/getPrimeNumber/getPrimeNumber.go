package main

import "fmt"

func main() {
	ch := generateNatural()
	for i := 0; i < 1000; i++ {
		prime := <-ch
		fmt.Printf("%v:%v\n", i+1, prime)
		ch = primeFilter(ch, prime)
	}
}

func generateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func primeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
