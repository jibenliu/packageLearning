package main

import (
	"curl/util"
	"fmt"
	"net/http"
)

func main() {
	checkUrl := "https://www.baidu.com"
	header := make(map[string]string)
	res, err := util.Hpool.Request(checkUrl, http.MethodGet, "", header)
	if err != nil {
		fmt.Printf("error is %v", err)
		return
	}
	fmt.Println(res)
}
