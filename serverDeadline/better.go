package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func slowAPICall() string {
	d := rand.Intn(5)
	select {
	case <-time.After(time.Duration(d) * time.Second):
		log.Printf("Slow API call done after %s seconds.\n", d)
		return "foobar"
	}
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	result := slowAPICall()
	io.WriteString(w, result+"\n")
}

func main() {
	srv := http.Server{
		Addr:         ":8888",
		WriteTimeout: 5 * time.Second,
		Handler:      http.TimeoutHandler(http.HandlerFunc(slowHandler), 1*time.Second, "Timeout!\n"),
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

