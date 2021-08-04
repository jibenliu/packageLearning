package main

import (
	"fmt"
	"html/template"
	"time"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello,", g.Name)
}

type EmailJob struct {
	Email string
}

func (e EmailJob) Run() {
	fmt.Println("Send,", e.Email)
}

func main() {
	r := gin.Default()

	jobrunner.Start()
	jobrunner.Every(5*time.Second, GreetingJob{Name: "dj"})
	jobrunner.Every(10*time.Second, EmailJob{Email: "935653229@qq.com"})

	r.GET("/jobrunner/json", JobJson)
	r.GET("/jobrunner/html", JobHtml)

	r.Run(":8888")
}

func JobJson(c *gin.Context) {
	c.JSON(200, jobrunner.StatusJson())
}

func JobHtml(c *gin.Context) {
	t, err := template.ParseFiles("./Status.html")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, "error")
	}
	t.Execute(c.Writer, jobrunner.StatusPage())
}


// http://localhost:8888/jobrunner/html

// 我们还可以通过http://localhost:8888/jobrunner/json获取原始 JSON 格式的数据自己处理：