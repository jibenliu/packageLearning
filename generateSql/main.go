package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("/Users/jibenliu/Downloads/vat_invoice_rule_template_project.sql")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	newFile, err := os.Create("/Users/jibenliu/Downloads/vat_invoice_rule_template_project-bak.sql")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer newFile.Close()
	buf := bufio.NewReader(f)
	line := 0
	for {
		b, errR := buf.ReadBytes('\n')
		if errR != nil && errR != io.EOF {
			fmt.Println(errR.Error())
			break
		}
		str := string(b)

		//match, _ := regexp.MatchString(`^\s*$`, str) //去掉空白行
		//if match && errR != io.EOF {
		//	continue
		//}
		//str = strings.TrimSpace(str) //去除行的左右两侧空白字符
		//if str == "" {
		//	continue
		//}

		//str = strings.TrimPrefix(str, "#") //去除以#开头的注释行
		//if str == "" {
		//	continue
		//}

		//index := strings.Index(str, "INSERT INTO")
		//if index >= 0 {
		//	line++
		//	m := line % 1000
		//	if m > 0 && line != 1 {
		//		str = strings.Replace(str, "INSERT INTO `employee` VALUES ", " ", 1)
		//	}
		//	if m != 999 && errR != io.EOF {
		//		str = strings.Replace(str, ";", ",", 1)
		//	}
		//}
		//_, _ = newFile.WriteString(str)
		//if errR == io.EOF {
		//	break
		//}

		index := strings.Index(str, "INSERT INTO")
		if index >= 0 {
			line++
			re3, _ := regexp.Compile("\\([0-9]*")
			str = re3.ReplaceAllString(str, "("+strconv.Itoa(line))
		}
		_, _ = newFile.WriteString(str)
		if errR == io.EOF {
			break
		}
	}
}
