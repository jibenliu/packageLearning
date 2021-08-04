package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	_ = t.Execute(w, "Hello world")
}

func parseString(w http.ResponseWriter, r *http.Request) {
	tmpl := `<!DOCTYPE html> <html>
        <head>
            <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
            <title>Go Web Programming</title>
        </head>
        <body>
            {{ . }}
        </body> 
    </html>`

	t := template.New("tmpl.html")
	t.Parse(tmpl)
	t.Execute(w, "Hello World!")
}

func main() {
	http.HandleFunc("/template", process)
	http.HandleFunc("/string", parseString)
	http.ListenAndServe(":8088", nil)
}
