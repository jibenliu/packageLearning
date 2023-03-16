package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/telegraf", telegraf)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

func telegraf(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	jsonDecoder := json.NewDecoder(r.Body)
	_ = jsonDecoder.Decode(&body)
	fmt.Println(body)
}
