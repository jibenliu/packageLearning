package main

import (
	"fmt"
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"time"
)

type EmailJob struct {
	Name  string
	Email string
}

type User struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (j EmailJob) Run() {
	e := email.NewEmail()
	e.From = "hh <jibenliu@126.com>"
	e.To = []string{j.Email}
	e.Cc = []string{"isuben@foxmail.com"}
	e.Subject = "Welcome To Awesome-Web"
	e.Text = []byte(fmt.Sprintf(`
  Hello, %s
  Welcome Back
  `, j.Name))
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("jibenliu", "jibenliu@126.com", "KFLEVYAVKWKBISLZ", "smtp.126.com"))
	if err != nil {
		log.Fatal(err)
	}
}

func login(c *gin.Context) {
	var u User
	if c.ShouldBind(&u) == nil {
		c.String(200, "login success")
		jobrunner.In(5*time.Second, EmailJob{Name: u.Name, Email: u.Email})
	} else {
		c.String(404, "login failed")
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", login)
	r.Run(":8888")
}

// http://localhost:8888/login?name=dj&email=jibenleo@163.com