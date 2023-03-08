package main

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
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
	f.SetActiveSheet(4)
	rows, err := f.GetRows("prompts管理")
	if err != nil {
		panic("读取excel sheet 失败" + err.Error())
	}

	for i := 2; i < len(rows); i++ {
		// Get value from cell by given worksheet name and cell reference.
		useCase, err := f.GetCellValue("prompts管理", "B"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		var prompt InspirationPrompt
		prompt.UseCase = useCase
		prompt.DefaultLanguageCode = "EN"
		englishTemplate, err := f.GetCellValue("prompts管理", "C"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		prompt.DefaultTemplate = englishTemplate
		re := regexp2.MustCompile(`(?<=\{)([^\}]+)`, 0) //找到大括号开头，非大括号中间的参数
		params := regexp2FindAllString(re, englishTemplate)
		prompt.DefaultParams = params
		if db.Where("use_case = ? AND default_language_code >= ?", useCase, "EN").First(&prompt).Error == gorm.ErrRecordNotFound {
			db.Create(&prompt)
		}
		chineseTemplate, err := f.GetCellValue("prompts管理", "D"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		var cPrompt InspirationPrompt
		cPrompt.UseCase = useCase
		cPrompt.DefaultLanguageCode = "CH"
		cTemplate := strings.Replace(chineseTemplate, "（", "(", -1)
		cTemplate = strings.Replace(cTemplate, "）", ")", -1)
		cPrompt.DefaultTemplate = cTemplate
		cRe := regexp2.MustCompile(`(?<=\{)([^\}]+)`, 0) //找到大括号开头，非大括号中间的参数
		cParams := regexp2FindAllString(cRe, cTemplate)
		cPrompt.DefaultParams = cParams
		if db.Where("use_case = ? AND default_language_code >= ?", useCase, "CH").First(&cPrompt).Error == gorm.ErrRecordNotFound {
			db.Create(&cPrompt)
		}
	}
}

func regexp2FindAllString(re *regexp2.Regexp, s string) string {
	var matches string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		if matches == "" {
			matches += m.String()
		} else {
			matches = matches + "," + m.String()
		}
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

// InspirationPrompt 提示类模板
type InspirationPrompt struct {
	gorm.Model
	UseCase             string  `gorm:"column:use_case" db:"use_case" json:"use_case" form:"use_case"`                                                     //创作类型
	DefaultLanguageCode string  `gorm:"column:default_language_code" db:"default_language_code" json:"default_language_code" form:"default_language_code"` //用户语言
	DefaultTitle        string  `gorm:"column:default_title" db:"default_title" json:"default_title" form:"default_title"`                                 //默认抬头
	DefaultDescription  string  `gorm:"column:default_description" db:"default_description" json:"default_description" form:"default_description"`         //默认描述
	DefaultTemperature  float64 `gorm:"column:default_temperature" db:"default_temperature" json:"default_temperature" form:"default_temperature"`         //默认随机度
	DefaultTemplate     string  `gorm:"column:default_template" db:"default_template" json:"default_template" form:"default_template"`                     //默认引导的模板
	DefaultText         string  `gorm:"column:default_text" db:"default_text" json:"default_text" form:"default_text"`                                     //默认文本
	DefaultParams       string  `gorm:"column:default_params" db:"default_params" json:"default_params" form:"default_params"`                             //默认参数
	DefaultLength       int     `gorm:"column:default_length" db:"default_length" json:"default_length" form:"default_length"`                             //默认回答的长度
}
