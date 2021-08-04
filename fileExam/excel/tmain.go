package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var GameListMap map[string]*GolandLtd

func init() {
	GameListMap = make(map[string]*GolandLtd)
	if !ReadCSV() {
		fmt.Println("初始化失败！")
	}
}

func main() {

}

type GolandLtd struct {
	ID       string
	GameName string
	GameType string
	GameDev  string
}

//读取本地的CSV
func ReadCSV() bool {
	fmt.Println("Golang语言社区 应用实例")
	fileName := "gameList.xlsx"
	fileName = "./csv/" + fileName

	//fileBytes, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	return false
	//}
	//r2 := csv.NewReader(strings.NewReader(string(fileBytes)))
	rfile, _ := os.Open(fileName)
	r := csv.NewReader(rfile)
	strs, _ := r.Read()
	fmt.Println(strs)
	//r2.Comma = ','
	//for {
	//	record, err := r2.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(record)
	//}
	//data, err := r2.ReadAll()
	//fmt.Println(err)
	//if err != nil {
	//	return false
	//}
	//fmt.Println("获取的数据：", data)
	//csvLen := len(data)
	//for i := 1; i < csvLen; i++ {
	//	DataTmp := new(GolandLtd)
	//	DataTmp.ID = data[i][0]
	//	DataTmp.GameName = data[i][1]
	//	DataTmp.GameType = data[i][2]
	//	DataTmp.GameDev = data[i][3]
	//	GameListMap[DataTmp.ID] = DataTmp
	//}
	//fmt.Println("获取的数据：", GameListMap)
	return true
}
