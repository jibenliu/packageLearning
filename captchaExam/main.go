package main

import (
	"fmt"
	"github.com/dchest/captcha"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/captcha/generate", func(w http.ResponseWriter, r *http.Request) {
		id := captcha.NewLen(6)
		if _, err := fmt.Fprintf(w, id); err != nil {
			log.Println("generate captcha error", err)
		}
	})

	http.HandleFunc("/captcha/image", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		if err := captcha.WriteImage(w, id, 240, 240); err != nil {
			log.Println("show captcha error", err)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Println("parseForm error", err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		// 获取验证码 ID 和验证码值
		id := r.FormValue("id")
		value := r.FormValue("value")
		// 比对提交的验证码值和内存中的验证码值
		if captcha.VerifyString(id, value) {
			_, _ = fmt.Fprint(w, "ok")
		} else {
			_, _ = fmt.Fprint(w, "mismatch")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
