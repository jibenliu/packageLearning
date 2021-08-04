package main

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func main() {
	e := email.NewEmail()
	e.From = "hh <jibenliu@126.com>"
	e.To = []string{"isuben@foxmail.com", "jibenleo@163.com"}
	//e.Cc = []string{"test1@126.com", "test2@126.com"} //抄送
	//e.Bcc = []string{"secret@126.com"} //秘密抄送
	e.Subject = "Awesome email"
	//e.Text = []byte("Text body is , of course, supported")
	e.HTML = []byte(`
<ul>
<li><a "https://darjun.github.io/2020/01/10/godailylib/flag/">Go 每日一库之 flag</a></li>
<li><a "https://darjun.github.io/2020/01/10/godailylib/go-flags/">Go 每日一库之 go-flags</a></li>
<li><a "https://darjun.github.io/2020/01/14/godailylib/go-homedir/">Go 每日一库之 go-homedir</a></li>
<li><a "https://darjun.github.io/2020/01/15/godailylib/go-ini/">Go 每日一库之 go-ini</a></li>
<li><a "https://darjun.github.io/2020/01/17/godailylib/cobra/">Go 每日一库之 cobra</a></li>
<li><a "https://darjun.github.io/2020/01/18/godailylib/viper/">Go 每日一库之 viper</a></li>
<li><a "https://darjun.github.io/2020/01/19/godailylib/fsnotify/">Go 每日一库之 fsnotify</a></li>
<li><a "https://darjun.github.io/2020/01/20/godailylib/cast/">Go 每日一库之 cast</a></li>
</ul>
`)
	e.AttachFile("main.go")
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("jibenliu", "jibenliu@126.com", "KFLEVYAVKWKBISLZ", "smtp.126.com"))
	if err != nil {
		log.Fatal(err)
	}
}
