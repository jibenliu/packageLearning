package server

import (
	"fmt"
	"net/http"
)

func ExampleServe() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request from client and host :", r.Host)
		fmt.Fprintln(w, r.Body)
	})
	http.ListenAndServe(":8080", nil)
}
