package main

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
	"uploadFiles/utils"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "chapter01/form_upload.html", nil)
	})

	router.POST("/", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		file, err := ctx.FormFile("file")
		if err != nil {
			fmt.Println("获取数据失败")
			ctx.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "获取数据失败",
			})
		} else {
			fmt.Println("接受的数据", name, file.Filename)
			//获取文件名称
			fmt.Println(file.Filename)
			//文件大小
			fmt.Println(file.Size)
			//获取文件的后缀名
			extString := path.Ext(file.Filename)
			fmt.Println(extString)
			//允许上传文件的格式
			allowExtMap := map[string]bool{
				".jpg":  true,
				".png":  true,
				".gif":  true,
				".jpeg": true,
			}
			if _, ok := allowExtMap[extString]; !ok {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code":    0,
					"message": "上传文件格式不支持",
				})
			}
			//根据当前时间鹾生成一个新的文件名
			fileNameInt := time.Now().Unix()
			fileNameStr := strconv.FormatInt(fileNameInt, 10)
			//新的文件名
			fileName := fileNameStr + extString
			//保存上传文件
			pathName, _ := utils.Mkdir("upload")
			filePath := filepath.Join(pathName, "/", fileName)
			//_ = ctx.SaveUploadedFile(file, filePath)
			client, err := oss.New("http://oss-cn-shenzhen.aliyuncs.com", "LTAI4Ff9jV7DfiPrJT36a", "zZOpRqGtKNQl30Su6Ytj12b3IEF")
			if err != nil {
				fmt.Println("阿里云上传错误", err)
				return
			}
			//指定存储空间
			bucket, err := client.Bucket("shuiping-code")
			if err != nil {
				fmt.Println("存储空间错误")
				os.Exit(-1)
			}
			//打开文件
			fileHandle, err := file.Open()
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    1,
					"message": "打开文件错误",
				})
				return
			}
			defer fileHandle.Close()
			fileByte, _ := ioutil.ReadAll(fileHandle)
			//上传到oss上
			err = bucket.PutObject(filePath, bytes.NewReader(fileByte))
			if err != nil {
				fmt.Println(err)
				ctx.JSON(http.StatusOK, gin.H{
					"code":    0,
					"message": "解析错误",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "success",
			})
		}
	})
	_ = router.Run()
}
