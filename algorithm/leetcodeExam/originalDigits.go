package main

import (
	"strconv"
	"strings"
)

/**
给定一个非空字符串，其中包含字母顺序打乱的英文单词表示的数字0-9。按升序输出原始的数字。

注意:

输入只包含小写英文字母。输入保证合法并可以转换为原始的数字，这意味着像 "abc" 或 "zerone" 的输入是不允许的。输入字符串的长度小于 50,000。示例 1:

输入: "owoztneoer"

输出: "012" (zeroonetwo)
示例 2:

输入: "fviefuro"

输出: "45" (fourfive)
提示：

1 <= s.length <= 105
s[i] 为 ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"] 这些字符之一
s 保证是一个符合题目要求的字符串

*/
func originalDigits(s string) string {
	count := make([]int, 26+'a') // a-z 数量为26，+a为最小值
	for _, b := range s {
		count[b]++
	}

	out := make([]int, 10)
	out[0] = count['z']
	out[2] = count['w']
	out[4] = count['u']
	out[6] = count['x']
	out[8] = count['g']
	out[1] = count['o'] - out[0] - out[2] - out[4]
	out[3] = count['h'] - out[8]
	out[5] = count['f'] - out[4]
	out[7] = count['s'] - out[6]
	out[9] = count['i'] - out[5] - out[6] - out[8]

	//res := ""
	var res strings.Builder // 效率更高
	for i, v := range out {
		for j := 0; j < v; j++ {
			//res += strconv.Itoa(i)
			res.WriteString(strconv.Itoa(i))
		}
	}

	return res.String()
}
