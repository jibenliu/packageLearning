package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
)

var url = "https://cdimage.uniontech.com/daily-iso/release/uniontechos-desktop-20-professional-1050-amd64-20220110-2053.iso"

var wg sync.WaitGroup

func main() {
	res, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK && res.Header.Get("Accept-Ranges") == "bytes" && res.ContentLength > 0 {
		//fmt.Println(res.ContentLength)
		//fmt.Println(res.ContentLength / int64(runtime.NumCPU()))
		rangeSize := math.Ceil(float64(res.ContentLength / int64(runtime.NumCPU())))
		//fmt.Println(rangeSize)
		wg.Add(runtime.NumCPU())
		for i := 0; i < runtime.NumCPU(); i++ {
			//fmt.Println((i+1)*int(rangeSize), "----", res.ContentLength)
			//continue
			rangeEnd := (i + 1) * int(rangeSize)
			if (i+1)*int(rangeSize) > int(res.ContentLength) {
				rangeEnd = int(res.ContentLength)
			}
			go download(url, i*int(rangeSize), rangeEnd, i)
		}
		wg.Wait()
		_, fi := filepath.Split(url)
		des, err := os.OpenFile("../cache/"+fi, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer des.Close()
		for i := 0; i < runtime.NumCPU(); i++ {
			source, err := os.OpenFile("../cache/"+fi+"."+strconv.Itoa(i), os.O_RDONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(des, source)
			source.Close()
			os.Remove("../cache/" + fi + "." + strconv.Itoa(i))
		}
	}
}

func download(url string, start, end, index int) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	res, err := http.DefaultClient.Do(req)
	_, fi := filepath.Split(url)
	f, err := os.OpenFile("../cache/"+fi+"."+strconv.Itoa(index), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()
}

//func Timeout(ctx context.Context){
//	for {
//		select {
//		case <- ctx.Done():
//			// user canceled the download
//			return
//		default:
//			_, err = io.CopyN(f, res.Body, BUFFER_SIZE))
//			if err != nil {
//				if err == io.EOF {
//					return
//				} else {
//					log.Fatal(err)
//				}
//			}
//		}
//	}
//}
