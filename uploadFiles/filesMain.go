package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
	"uploadFiles/utils"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "chapter01/form_upload2.html", nil)
	})
	router.POST("/", func(ctx *gin.Context) {
		if form, err := ctx.MultipartForm(); err == nil {
			//1.获取文件
			files := form.File["file"]
			//2.循环全部的文件
			for _, file := range files {
				// 3.根据时间鹾生成文件名
				fileNameInt := time.Now().Unix()
				fileNameStr := strconv.FormatInt(fileNameInt, 10)
				//4.新的文件名(如果是同时上传多张图片的时候就会同名，因此这里使用时间鹾加文件名方式)
				fileName := fileNameStr + file.Filename
				//5.保存上传文件
				pathName, _ := utils.Mkdir("upload")
				filePath := filepath.Join(pathName, "/", fileName)
				_ = ctx.SaveUploadedFile(file, filePath)
			}
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "上传成功",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "获取数据失败",
			})
		}
	})
	_ = router.Run()
}
