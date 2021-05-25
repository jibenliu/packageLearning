package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("current time is %v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minutes := now.Minute()
	seconds := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minutes, seconds) //%02d是指宽度，不够就补0

	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Printf("现在的时间戳：%v\n", timeStamp1)
	fmt.Printf("现在的纳秒时间戳：%v\n", timeStamp2)

	timeObj := time.Unix(timeStamp1, 0)
	fmt.Printf("timeObj:%v\n", timeObj)

	t := time.Now()
	fmt.Printf("今天是%v\n", t.Weekday().String())

	later := now.Add(time.Hour)
	fmt.Printf("一个小时后的时间是%v\n", later)

	//定时器
	//ticker := time.Tick(time.Second * 5) //定义一个5秒间隔的定时器
	//for i := range ticker {
	//	fmt.Println(i)
	//}

	//时间格式化
	// 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-15-4-5
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("15:04:05 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("2006-01-02 13:05:01"))

}
