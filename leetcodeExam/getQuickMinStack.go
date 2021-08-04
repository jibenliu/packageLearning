package main

import "math"

/**
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
*/
type MinStack struct {
	data  []int //数据栈，存数据
	extra []int //辅助栈，存最小值
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{
		data:  []int{},
		extra: []int{math.MaxInt64},
	}
}

func (s *MinStack) Push(x int) {
	// 数据栈
	s.data = append(s.data, x)
	// 辅助栈
	s.extra = append(s.extra, min(x, s.extra[len(s.extra)-1]))
}

func (s *MinStack) Pop() {
	if len(s.data) != 0 {
		s.data = s.data[:len(s.data)-1]
		s.extra = s.extra[:len(s.extra)-1]
	}
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	return s.extra[len(s.extra)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
