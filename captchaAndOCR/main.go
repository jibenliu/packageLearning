package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"io/ioutil"
)

var client *gosseract.Client

func main() {
	client = gosseract.NewClient()
	defer client.Close()
	getFileList("./images/")
}

func getFileList(path string) {
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			getFileList(path + file.Name() + "/")
		} else {
			ocrImage(path + file.Name())
		}
	}
}

func ocrImage(path string) {
	client.SetImage(path)
	text, _ := client.Text()
	fmt.Println(text)
}
