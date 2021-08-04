package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func readline1(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)
	for {
		_, err := rd.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
	}
}

func readline2(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)
	for {
		_, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
	}
}
func readline3(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)
	for {
		_, err := rd.ReadSlice('\n')
		if err != nil || err == io.EOF {
			break
		}
	}
}
func readline4(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)
	for {
		_, _, err := rd.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
	}
}

func testfile2(filename string) {
	fmt.Printf("============test2 %s ===========\n", filename)
	start := time.Now()
	readline1(filename)
	t1 := time.Now()
	fmt.Printf("Readline 1 cost: %v\n", t1.Sub(start))
	readline2(filename)
	t2 := time.Now()
	fmt.Printf("Readline 2 cost: %v\n", t2.Sub(t1))
	readline3(filename)
	t3 := time.Now()
	fmt.Printf("Readline 3 cost: %v\n", t3.Sub(t2))
	readline4(filename)
	t4 := time.Now()
	fmt.Printf("Readline 4 cost: %v\n", t4.Sub(t3))
}

func main() {
	testfile2("small.txt")
	testfile2("medium.txt")
	testfile2("large.zip")
}

// 读取文件中一行内容时，ReadSlice和ReadLine性能优于ReadBytes和ReadString，但由于ReadLine对换行的处理更加全面（兼容\n和\r\n换行），因此，实际开发过程中，建议使用ReadLine函数。