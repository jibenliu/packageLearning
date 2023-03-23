package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/telegraf", func(c *gin.Context) {
		query := c.Request.URL.Query()
		var queryMap = make(map[string]any, len(query))
		for k := range queryMap {
			queryMap[k] = c.Query(k)
		}
		fmt.Println("query", queryMap)
		req := make(map[string]any)
		for k, v := range c.Request.PostForm {
			req[k] = v[0]
		}
		fmt.Println("form", req)
		jParam, _ := c.GetRawData()
		var jReq map[string]any
		_ = json.Unmarshal(jParam, &jReq)
		fmt.Println("json", jReq)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
