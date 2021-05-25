package main

import "testing"

func TestPushPop(t *testing.T){
	c := new(Stack)
	for i:=0;i<11 ;i++  {
		c.Push(i)
	}
	if c.Pop() != 5{
		t.Log("Pop doesn't give 5")
		t.Fail()
	}
}
