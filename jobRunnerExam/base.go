package main

import (
	"fmt"
	"github.com/bamzi/jobrunner"
	"time"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello, ", g.Name)
}

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 5s", GreetingJob{Name: "dj"})
	time.Sleep(20 * time.Second)
}
