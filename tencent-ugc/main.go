package main

import (
	"encoding/base64"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tms/v20201229"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"strconv"
	"strings"
	"tencent-ugc/tool"
	"time"
)

const SecretID = "XXXX"
const SecretKey = "XXXX"

func main() {
	tool.ReadJsonFile()
	client, request := ugcRequest()
	readExcel(client, request)
}

var rr = rand.New(rand.NewSource(time.Now().UnixNano()))

func getRandomContent() string {
	l := len(tool.QSlice)
	index := rr.Intn(l)
	tmp := tool.QSlice[index]
	ts := strings.Trim(fmt.Sprintf("%s", tmp["q_content"]), " ")
	if ts != "" {
		return ts
	} else {
		return getRandomContent()
	}
}

func insertQuota(in string, content string) string {
	// 将 GBK 编码的汉字字符串转换为 UTF-8编码的字节串
	runes := []rune(content)
	l := len(runes)
	index := rr.Intn(l)
	return string(runes[:index]) + in + string(runes[index:])
}

func ugcRequest() (*tms.Client, *tms.TextModerationRequest) {
	credential := common.NewCredential(SecretID, SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tms.tencentcloudapi.com"
	client, _ := tms.NewClient(credential, "ap-guangzhou", cpf)
	request := tms.NewTextModerationRequest()
	return client, request
}

func readExcel(client *tms.Client, request *tms.TextModerationRequest) {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	f.SetActiveSheet(1)
	rows, err := f.GetRows("中文")
	if err != nil {
		panic("读取excel sheet 失败" + err.Error())
	}

	for i := 2; i < len(rows); i++ {
		words, err := f.GetCellValue("中文", "B"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		sentence := getRandomContent()
		sentence = insertQuota(words, sentence)
		err = f.SetCellValue("中文", "C"+strconv.Itoa(i), sentence)
		if err != nil {
			panic(err)
		}
		request.Content = common.StringPtr(base64.StdEncoding.EncodeToString([]byte(sentence)))
		//request.BizType = common.StringPtr("test_tms_strategy_bak")
		response, err := client.TextModeration(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		b := "放行"
		if *response.Response.Suggestion == "Block" {
			b = "拦截"
		}
		if *response.Response.Suggestion == "Review" {
			b = "存疑"
		}
		err = f.SetCellValue("中文", "D"+strconv.Itoa(i), b)
		if err != nil {
			panic(err)
		}
		l := *response.Response.Label
		if label, ok := labelMap[*response.Response.Label]; ok {
			l = label
		}
		err = f.SetCellValue("中文", "E"+strconv.Itoa(i), l)
		if err != nil {
			panic(err)
		}
		err = f.SetCellValue("中文", "F"+strconv.Itoa(i), *response.Response.Score)
		if err != nil {
			panic(err)
		}
		f.Save()
	}
}

var labelMap = map[string]string{
	"Normal":  "正常",
	"Porn":    "色情",
	"Abuse":   "谩骂",
	"Ad":      "广告",
	"Polity":  "政治",
	"Terror":  "暴乱",
	"Illegal": "违法",
}
