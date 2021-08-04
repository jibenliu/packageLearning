package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"protocExam/protobuf/pro"
)

func write() {
	p1 := &pro.Person{ //实例对象要&
		Id:   1,
		Name: "张三",
		Phones: []*pro.Phone{ //数组前面要加*
			{
				Type:   pro.PhoneType_HOME, //赋值是key : value
				Number: "123",
			},
			{
				Type:   pro.PhoneType_WORK,
				Number: "321",
			},
		},
	}

	book := &pro.ContactBook{}
	book.Persons = append(book.Persons, p1)

	data, _ := proto.Marshal(book)

	ioutil.WriteFile("./test.txt", data, os.ModePerm)

}

func read() {

	data, _ := ioutil.ReadFile("./test.txt")
	book := &pro.ContactBook{}

	proto.Unmarshal(data, book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}
}

func main() {
	write()
	read()
}
