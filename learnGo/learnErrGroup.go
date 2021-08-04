package main

import (
	"fmt"
	"github.com/golang/sync"
	"net/http"
)

func main() {
	//var g sync.Group
	//var urls = []string{
	//	"http://www.golang.org/",
	//	"http://www.google.com/",
	//	"http://www.somestupidname.com/",
	//}
	//for i := range urls {
	//	url := urls[i]
	//	g.Go(func() error {
	//		resp, err := http.Get(url)
	//		if err == nil {
	//			resp.Body.Close()
	//		}
	//		return err
	//	})
	//}
	//if err := g.Wait(); err == nil {
	//	fmt.Println("Successfully fetched all URLs.")
	//} else {
	//	fmt.Printf("有异常%v", err)
	//}
}
