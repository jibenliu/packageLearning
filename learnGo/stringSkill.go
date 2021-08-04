package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	//多行字符串
	var str = `
		你好吗
		This is a test string
	`
	fmt.Println(str)

	//字符串拼接1
	var b bytes.Buffer
	for i := 0; i < 1000; i++ {
		b.WriteString(randString(1))
	}
	fmt.Println(b.String())
	//字符串拼接2
	var c []string
	for i := 0; i < 1000; i++ {
		c = append(c, randString(1))
	}
	fmt.Println(strings.Join(c, ""))
	//字符串拼接3
	var a strings.Builder
	for i := 0; i < 1000; i++ {
		a.WriteString("string3")
	}
	fmt.Println(a.String())
	//字符串拼接4
	bs := make([]byte, 1000)
	bl := 0
	for n := 0; n < 1000; n++ {
		bl += copy(bs[bl:], "string4") //用于将源slice的数据（第二个参数）,复制到目标slice（第一个参数） 返回值为拷贝了的数据个数,是len(dst)和len(src)中的最小值
	}
	fmt.Println(string(bs))
	//字符串拼接5
	var result string
	for i := 0; i < 1000; i++ {
		result += "string5"
	}
	fmt.Println(result)
	//字符串拼接6
	fmt.Println(strings.Repeat("x6",1000))
	return

	//任意类型转字符串
	var d = 123
	s := string(d)
	fmt.Println(s)
	fmt.Println(strconv.Itoa(d))

	//随机字符串
	var e []string
	for i := 0; i < 1000; i++ {
		e = append(e, RandString(1))
	}
	fmt.Println(strings.Join(e, ""))
}

var s = rand.NewSource(time.Now().UnixNano())

func randString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	stringBytes := []byte(str)
	var result []byte
	r := rand.New(s)
	for i := 0; i < length; i++ {
		result = append(result, stringBytes[r.Intn(len(stringBytes))])
	}
	return string(result)
}

var source = rand.NewSource(time.Now().UnixNano())

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}
