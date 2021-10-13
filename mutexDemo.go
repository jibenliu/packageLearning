package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//type URL struct {
//	Ip     string
//	Port   string
//	mux    sync.RWMutex
//	params url.Values
//}
//
//func (c *URL) Clone() URL {
//	newUrl := URL{}
//	newUrl.Ip = c.Ip
//	newUrl.params = url.Values{}
//	return newUrl
//}

var random = rand.New(rand.NewSource(time.Now().Unix()))
var kinds = []int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122}

func main() {
	//c := URL{}
	//fmt.Println(c.Clone())
	var result sync.Map
	var wg sync.WaitGroup
	goroutines := 20
	wg.Add(goroutines)
	for s := 0; s < goroutines; s++ {
		go func(s int) {
			for i := 50 * (s - 1); i < 5*s; i++ {
				//var random = rand.New(rand.NewSource(time.Now().Unix()))
				//time.Sleep(time.Millisecond * 200)
				/**
				每次操作的间隔都在毫秒级下，所以每次通过time.Now().Unix()取出来的时间戳都是同一个值,即使用了同一个seed
				每次rand都会使用相同的seed来生成随机队列，这样一来在循环中使用相同seed得到的随机队列都是相同的，而生成随机数时每次都会去取同一个位置的数，所以每次取到的随机数都是相同的
				seed 只用于决定一个确定的随机序列。不管seed多大多小，只要随机序列一确定，本身就不会再重复。除非是样本空间太小
				 */
				result.Store(i, string(rune(kinds[random.Intn(62)])))
			}
			wg.Done()
		}(s)
	}
	wg.Wait()
	result.Range(func(key, value interface{}) bool {
		fmt.Println(value)
		return true
	})
}
