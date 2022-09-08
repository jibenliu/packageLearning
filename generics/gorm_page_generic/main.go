package main

import (
	"gorm_page_generic/database"
	"gorm_page_generic/model"
	"gorm_page_generic/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	cityService    service.CityService
	countryService service.CountryService
)

func main() {
	if err := database.InitMysql(); err != nil {
		log.Fatalln("数据库连接出错")
	}
	defer database.Close()
	r := gin.Default()
	r.POST("/city/list", func(c *gin.Context) {
		var queryVo model.CityQueryInfo
		if e := c.ShouldBindJSON(&queryVo); e != nil {
			c.JSON(http.StatusOK, gin.H{"code": 300, "msg": "参数错误"})
			return
		}
		pageResponse, e := cityService.SelectPageList(queryVo)
		if e != nil {
			c.JSON(http.StatusOK, model.Response{Code: 400, Msg: "操作失败"})
			return
		}
		c.JSON(
			http.StatusOK,
			model.Response{Code: 200, Msg: "操做成功", Data: pageResponse},
		)
	})

	r.POST("/country/list", func(c *gin.Context) {
		var queryVo model.CountryQueryInfo
		if e := c.ShouldBindJSON(&queryVo); e != nil {
			c.JSON(http.StatusOK, model.Response{Code: 300, Msg: "参数错误"})
			return
		}
		pageResponse, e := countryService.SelectPageList(queryVo)
		if e != nil {
			c.JSON(http.StatusOK, model.Response{Code: 400, Msg: "操作失败"})
			return
		}
		c.JSON(
			http.StatusOK,
			model.Response{Code: 200, Msg: "操做成功", Data: pageResponse},
		)
	})
	r.Run(":8080")
}
