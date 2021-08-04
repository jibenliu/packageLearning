package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// 读取文件大小 自带的read
func read1(filename string) int {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 4096)
	var nBytes int
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		nBytes += n
	}
	return nBytes
}

// bufio来读取
func read2(filename string) int {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 4096)
	var nBytes int
	rd := bufio.NewReader(fi)
	for {
		n, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		nBytes += n
	}
	return nBytes
}

// ioutil读取
func read3(filename string) int {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	nBytes := len(fd)
	return nBytes
}

func read4(filename string) int {
	fi, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
	return int(fi.Size())
}

func testfile1(filename string) {
	fmt.Printf("============test1 %s ===========\n", filename)
	start := time.Now()
	size1 := read1(filename)
	t1 := time.Now()
	fmt.Printf("Read 1 cost: %v, size: %d\n", t1.Sub(start), size1)
	size2 := read2(filename)
	t2 := time.Now()
	fmt.Printf("Read 2 cost: %v, size: %d\n", t2.Sub(t1), size2)
	size3 := read3(filename)
	t3 := time.Now()
	fmt.Printf("Read 3 cost: %v, size: %d\n", t3.Sub(t2), size3)
	size4 := read4(filename)
	t4 := time.Now()
	fmt.Printf("Read 4 cost: %v, size: %d\n", t4.Sub(t3), size4)
}

func main() {
	testfile1("small.txt")
	testfile1("medium.txt")
	testfile1("large.zip")
}

// 当读取小文件时，使用ioutil效率明显优于os和bufio，但如果是大文件，bufio读取会更快。
