package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := excelize.OpenFile("doc.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	f.SetActiveSheet(1)
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		panic("读取excel sheet 失败" + err.Error())
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var min, max = 1, 6
	for i := 2; i < len(rows); i++ {
		id, err := f.GetCellValue("Sheet1", "A"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		nickname, err := f.GetCellValue("Sheet1", "B"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		ask, err := f.GetCellValue("Sheet1", "C"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}

		answer, err := f.GetCellValue("Sheet1", "D"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		avatar := r1.Intn(max-min) + min
		var content []map[string]string
		content = append(content, map[string]string{"ask": strings.TrimSpace(ask), "answer": strings.TrimSpace(answer)})
		cj, _ := json.Marshal(content)
		sql := "INSERT INTO `ai_chat_shared_questions`(`question_id`,`question_start`,`content`,`nickname`,`avatar`)VALUES(" + strings.TrimSpace(id) + "," + strconv.Quote(strings.TrimSpace(ask)) + "," + strconv.Quote(string(cj)) + "," + strconv.Quote(strings.TrimSpace(nickname)) + "," + strconv.Itoa(avatar) + ");"
		fmt.Println(sql)
	}
}
