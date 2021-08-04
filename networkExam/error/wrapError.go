package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
)

type Error interface {
	Caller() []CallerInfo
	Wrapped() []error
	Code() int
	error
	private()
}

type CallerInfo struct {
	FuncName string
	FileName string
	FileLine string
}

// func New(msg string) error
//func NewWithCode(code int, msg string) error
//func Wrap(err error, msg string) error
//func WrapWithCode(code int, err error, msg string) error
//func FromJson(json string) (Error, error)
//func ToJson(err error) string
func loadConfig() error {
	_, err := ioutil.ReadFile("/path/to/file")
	if err != nil {
		return errors.Wrap(err, "read failed")
	}
	// ...
}

func setup() error {
	err := loadConfig()
	if err != nil {
		return errors.Wrap(err, "invalid config")
	}
	// ...
}

// 以JSON字符串方式发送错误
func sendError(ch chan<- string, err error) {
	ch <- errors.ToJson(err)
}

// 接收JSON字符串格式的错误
func recvError(ch <-chan string) error {
	p, err := errors.FromJson(<-ch)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
func main() {
	defer func() { //违背了go简单直接的编程哲学
		if r := recover(); r != nil {
			switch x := r.(type) {
			case runtime.Error:
				fmt.Println(x.Error())
			// 这是运行时错误类型异常
			case error:
			// 普通错误类型异常
			default:
				// 其他类型异常
				// ...
			}
		}
	}()
}
