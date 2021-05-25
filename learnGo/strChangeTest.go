package main

import (
	"fmt"
	"github.com/gotail/strx"
	_ "github.com/gotail/strx"
)

func main() {
	var newStr = strx.New("这*里|||有人民yy币  $yy#{money}   ..").
		NoConsecutiveSpaces().
		Trim(" ").
		Clean("*", "|", "yy").
		Replace("$", "￥").
		Replace("#{money}", "250").
		TrimRight(".").
		Append("元").
		Val()

	fmt.Println(newStr)
}
