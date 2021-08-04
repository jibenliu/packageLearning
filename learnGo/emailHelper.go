package main

import (
	"github.com/go-gomail/gomail"
	"strings"
)

type EmailParam struct {
	//ServerHost邮箱服务器地址
	ServerHost string
	//邮箱服务器端口
	ServerPort int
	//发件人邮箱地址
	FromEmail string
	//发件人邮箱密码
	FromPassword string
	//接受者
	Receivers string
	//抄送者
	ReceiverPartners string
}

var serverHost, fromEmail, fromPassword string
var serverPort int

var m *gomail.Message

func InitEmail(ep *EmailParam) {
	var receivers []string

	serverHost = ep.ServerHost
	serverPort = ep.ServerPort
	fromEmail = ep.FromEmail
	fromPassword = ep.FromPassword

	m = gomail.NewMessage()

	if len(ep.Receivers) == 0 {
		return
	}

	for _, tmp := range strings.Split(ep.Receivers, ",") {
		receivers = append(receivers, strings.TrimSpace(tmp))
	}

	m.SetHeader("To", receivers...)
	if len(ep.ReceiverPartners) != 0 {
		for _, tmp := range strings.Split(ep.ReceiverPartners, ",") {
			receivers = append(receivers, strings.TrimSpace(tmp))
		}
		m.SetHeader("Cc", receivers...)
	}

	m.SetAddressHeader("From", fromEmail, "")
}

func SendEmail(subject, body string) {
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(serverHost, serverPort, fromEmail, fromPassword)
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}

func main() {
	serverHost = "smtp.qq.com"
	serverPort := 465
	fromEmail := "xxxxx@qq.com"
	fromPassword := "xxxxxx"

	receiver := "11111@11.com"
	receiverPartner := ""

	subject := "这是邮件主题"
	body := `这是正文<br>
	Hello<a href="http://www.baidu.com/">百度</a>`

	myEmail := &EmailParam{
		ServerHost:       serverHost,
		ServerPort:       serverPort,
		FromEmail:        fromEmail,
		FromPassword:     fromPassword,
		Receivers:        receiver,
		ReceiverPartners: receiverPartner,
	}

	InitEmail(myEmail)
	SendEmail(subject, body)

}
