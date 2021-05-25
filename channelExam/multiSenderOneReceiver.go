package main

import (
	"context"
	"fmt"
	"sync"
)

var wgSender sync.WaitGroup
var wgReceiver sync.WaitGroup

type ProcessMessage func(ctx context.Context, originMsg *Person) (resMsg *Person)

type Person struct {
	Name string
	Job  string
}

func main() {
	ctx := context.Background()
	firstPerson := &Person{Name: "John", Job: "worker"}
	secondPerson := &Person{Name: "Tom", Job: "engineer"}
	originalMsg := []*Person{firstPerson, secondPerson}
	ch := make(chan *Person, len(originalMsg))

	StartSend(ctx, originalMsg, ch, process)

	wgReceiver.Add(1)
	var resList []*Person
	go func() {
		resList = StartReceive(ch)
	}()
	wgSender.Wait()
	close(ch)
	wgReceiver.Wait()
	fmt.Println(resList[0].Job, resList[1].Job)
}

func StartSend(ctx context.Context, originalMsgs []*Person, ch chan *Person, process ProcessMessage) {
	for _, originalMsg := range originalMsgs {
		wgSender.Add(1)
		go Sender(ctx, originalMsg, ch, process)
	}
}

func StartReceive(ch chan *Person) (resList []*Person) {
	defer wgReceiver.Done()
	resList = Receiver(ch)
	return
}

func Sender(ctx context.Context, originalMsg *Person, ch chan *Person, process ProcessMessage) {
	defer wgSender.Done()
	ch <- process(ctx, originalMsg)

}

func Receiver(ch chan *Person) (eleList []*Person) {
	for {
		select {
		case ele, isOk := <-ch:
			if !isOk {
				return eleList
			}
			eleList = append(eleList, ele)
		default:
		}

	}
}

func process(ctx context.Context, originMsg *Person) (resMsg *Person) {
	originMsg.Job = "good " + originMsg.Job
	return originMsg
}
