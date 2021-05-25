package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type chunk struct {
	bufSize int
	offset  int64
}

// 在一个50gb的文件里用4gb的内存计算一个单词的出现次数
func main() {
	const BufferSize = 1000
	file, err := os.Open("medium.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := int(fileInfo.Size())
	concurrency := fileSize / BufferSize
	if remainder := fileSize % BufferSize; remainder != 0 { //实现的是floor方法
		concurrency++
	}
	chunkSizes := make([]chunk, concurrency)

	for i := 0; i < concurrency; i++ {
		chunkSizes[i].bufSize = BufferSize
		chunkSizes[i].offset = int64(BufferSize * i)
	}

	if remainder := fileSize % BufferSize; remainder != 0 {
		c := chunk{bufSize: remainder, offset: int64(concurrency * BufferSize)}
		concurrency++
		chunkSizes = append(chunkSizes, c)
	}

	channelString := make(chan []string, concurrency)

	for i := 0; i < concurrency; i++ {
		go func(chunkSizes []chunk, i int) {
			chunk := chunkSizes[i]
			buffer := make([]byte, chunk.bufSize)
			_, err := file.ReadAt(buffer, chunk.offset) //第一个参数是读取字节数
			if err != nil && err != io.EOF {
				fmt.Println(err)
				return
			}

			scanner := bufio.NewScanner(strings.NewReader(string(buffer)))
			scanner.Split(bufio.ScanWords) // 根据空格切分字符
			var temp []string
			for scanner.Scan() {
				temp = append(temp, scanner.Text())
			}

			channelString <- temp
		}(chunkSizes, i)
	}

	dictSorter := DictSorter{}
	ch := make([]string, 1)

	for i := 0; i < concurrency; i++ {
		ch = <-channelString
		for _, temp := range ch {
			exists := false
			for index, dict := range dictSorter.dicts {
				if dict.key == temp {
					exists = true
					dictSorter.dicts[index].value++
					break
				}
			}
			if exists == false {
				dictSorter.dicts = append(dictSorter.dicts, Dict{
					key:   temp,
					value: 1,
				})
			}
		}
	}
	sort.Sort(&dictSorter)
	fmt.Println(dictSorter.Len())
	for _, dict := range dictSorter.dicts {
		fmt.Println(dict)
	}
}

type Dict struct {
	key   string
	value int
}

type DictSorter struct {
	dicts []Dict
}

func (d DictSorter) by(p1, p2 *Dict) bool {
	return p1.value > p2.value
}

func (d *DictSorter) Len() int {
	return len(d.dicts)
}

func (d *DictSorter) Swap(i, j int) {
	d.dicts[i], d.dicts[j] = d.dicts[j], d.dicts[i]
}

func (d *DictSorter) Less(i, j int) bool {
	return d.by(&d.dicts[i], &d.dicts[j])
}
