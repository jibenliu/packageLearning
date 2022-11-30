package main

import (
	"bufio"
	"encoding/base64"
	client2 "faceFushion/client"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const SecretID = ""
const SecretKey = ""
const FaceFusionEndPoint = "facefusion.tencentcloudapi.com"
const AgeChangerEndPoint = "ft.tencentcloudapi.com"
const ImageSavePath = "./data/generate/"
const ProjectID = "314596"

var BabyFaceModelMap = map[string]string{
	"qc_314596_195673_4":  "baby_1y_boy_1",
	"qc_314596_591250_5":  "baby_1y_boy_2",
	"qc_314596_944170_6":  "baby_1y_boy_3",
	"qc_314596_217841_7":  "baby_1y_girl_1",
	"qc_314596_587972_8":  "baby_1y_girl_2",
	"qc_314596_590706_9":  "baby_1y_girl_3",
	"qc_314596_787200_10": "baby_3y_boy_1",
	"qc_314596_186391_16": "baby_3y_boy_2",
	"qc_314596_954800_17": "baby_3y_boy_3",
	"qc_314596_995740_18": "baby_3y_girl_1",
}

func main() {
	now := time.Now()
	logDir := "./data/male"
	var files []string
	err := filepath.Walk(logDir, func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, ".DS_Store") && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, fi := range files {
		for modelId, templateName := range BabyFaceModelMap {
			url := younger(fi)
			ret := faceFusion([]string{url}, modelId)
			saveImage(ret, templateName+"_"+strings.TrimSuffix(filepath.Base(fi), filepath.Ext(fi)))
		}
	}

	period := time.Since(now)
	fmt.Println("全部耗时时间: ", period)
}

func getClient(endPoint string) *client2.Client {
	credential := common.NewCredential(SecretID, SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 30
	cpf.HttpProfile.Endpoint = endPoint
	cpf.SignMethod = "TC3-HMAC-SHA256"
	client, _ := client2.NewClient(credential, regions.Guangzhou, cpf)
	return client
}

func faceFusion(urls []string, modelId string) string {
	client := getClient(FaceFusionEndPoint)
	request := client2.NewFuseFaceRequest()
	request.ProjectId = common.StringPtr(ProjectID)
	request.ModelId = common.StringPtr(modelId)

	for _, url := range urls {
		request.MergeInfos = append(request.MergeInfos, &client2.MergeInfo{Url: common.StringPtr(url)})
	}
	request.RspImgType = common.StringPtr("url")

	now := time.Now()
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.FuseFace(request)
	period := time.Since(now)
	fmt.Println("合成间隔时间: ", period)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return ""
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf("%s\n", response.ToJsonString())
	return *response.Response.FusedImage
}

func younger(filePath string) string {
	client := getClient(AgeChangerEndPoint)
	request := client2.NewChangeAgePicRequest()
	request.AgeInfos = []*client2.AgeInfo{
		{Age: common.Int64Ptr(10)},
	}
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	request.Image = common.StringPtr(base64.StdEncoding.EncodeToString(data))
	request.RspImgType = common.StringPtr("url")
	now := time.Now()
	response, err := client.ChangeAgePic(request)
	period := time.Since(now)
	fmt.Println("年龄变幻间隔时间: ", period)
	if err != nil {
		panic(err)
	}
	return *response.Response.ResultUrl
}

func saveImage(imgUrl string, fileName string) {
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(ImageSavePath + fileName + ".png")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}
