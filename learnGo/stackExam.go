package main

//实现一个长度为10的栈
import (
	"fmt"
	"strconv"
)

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}

func (s *Stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
func main() {
	var s Stack
	s.Push(25)
	s.Push(14)
	fmt.Printf("stack %v\n", s)
	fmt.Printf(s.String())
}
