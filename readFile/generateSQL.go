package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//func pathExists(path string) (bool, error) {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//}

func processLine(line []byte) {
	newFilePath := "new_vat_tax_content_chart.sql"
	//fileExist, err := pathExists(newFilePath)
	//if err != nil {
	//	return
	//}
	//if !fileExist {
	//	f, err := os.Create("new_vat_tax_content_chart.sql")
	//	if err != nil {
	//		return
	//	}
	//	defer f.Close()
	//}
	f, er := os.OpenFile(newFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if er != nil {
		return
	}
	defer f.Close()
	var lineString = string(line[:])
	var temp = strings.Replace(lineString, "VALUES (", "SELECT ", 1)
	temp = strings.Replace(temp, ");", " FROM DUAL WHERE NOT EXISTS(SELECT * FROM `vat_tax_content_chart` WHERE id = "+strconv.FormatInt(index, 10)+");\n", 1)
	_, newErr := f.WriteString(fmt.Sprintf(temp))
	if newErr != nil {
		fmt.Println(newErr)
		return
	}
	//_, _ = os.Stdout.Write(line)
}

func ReadLine(filePath string, hookFn func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookFn(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		index++
	}
}

var index int64

func main() {
	index = -33
	_ = ReadLine("./vat_tax_content_chart.sql", processLine)
}
