package main

import (
	"fmt"
	"rsc.io/pdf"
)

func main() {
	reader, err := pdf.Open("ntf.pdf")
	if err != nil {
		panic(err)
	}
	page := reader.NumPage()
	fmt.Println("pdf 总页数", page)

	p := reader.Page(11)
	c := p.Content()
	ss := ""
	for _, s := range c.Text {
		ss += s.S
	}
	fmt.Println(ss)
}
