package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	path := "../learnBeego/controllers"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		filePath := path + "\\" + strings.Replace(f.Name(), "vat_", "", 1)
		_ = os.Rename(path+"\\"+f.Name(), filePath)
		fmt.Println(filePath)
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModeAppend)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			continue
		}
		bfRd := bufio.NewReader(f)
		index := 0
		for {
			lineString, err := bfRd.ReadBytes('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			var temp = strings.Replace(lineString, "Vat", "", 1)
			fmt.Println(temp)
			_, _ = f.Seek(int64(-(len(lineString))), 1)
			_, newErr := f.WriteString(fmt.Sprintf(temp) + "\n")
			if newErr != nil {
				fmt.Println(newErr)
				break
			}
			index++
		}
		f.Close()
		break
	}
}
