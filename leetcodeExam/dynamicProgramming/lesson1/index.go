package main

import (
	"fmt"
	"runtime"
)

func main() {
	//f := bufio.NewReader(os.Stdin) //读取输入的内容
	//fmt.Print("输入你要运算的n:")
	//var input string
	//input, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
	var n int
	fmt.Printf("输入你要运算的n:")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Printf("输入错误，错误信息为：%v", err)
	}
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()

	fmt.Println("总方法数为：", climbStairs(n))

}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
