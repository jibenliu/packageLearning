package main

import (
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
	"os"
)

//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var store = sessions.NewFilesystemStore("sessionCache", []byte(os.Getenv("SESSION_KEY)")))

func counter(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "SESSION_ID_KEY")
	count := session.Values["count"]
	if count == nil {
		session.Values["count"] = 1
	} else {
		session.Values["count"] = count.(int) + 1
	}
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, _ := template.ParseFiles("counter.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, session.Values["count"])
}

func main() {
	http.HandleFunc("/counter", counter)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
