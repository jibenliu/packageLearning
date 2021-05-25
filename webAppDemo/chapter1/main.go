package main

import (
	"fmt"
	"net/http"
	"newTest/webAppDemo/chapter1/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "fresh 命令 热加载 URL.Path = %q\n", r.URL.Path)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
