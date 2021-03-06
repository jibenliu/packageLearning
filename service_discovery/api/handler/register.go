package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/service_discovery/configs"
	"github.com/skyhackvip/service_discovery/global"
	"github.com/skyhackvip/service_discovery/model"
	"github.com/skyhackvip/service_discovery/pkg/errcode"
	"log"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	log.Println("request api/register...")
	var req model.RequestRegister
	if e := c.ShouldBindJSON(&req); e != nil {
		log.Println("error:", e)
		err := errcode.ParamError
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"message": err.Error(),
		})
		return
	}
	//bind instance
	instance := model.NewInstance(&req)
	if instance.Status == 0 || instance.Status > 2 {
		log.Println("register params status invalid")
		err := errcode.ParamError
		c.JSON(http.StatusOK, gin.H{
			"code":    err.Code(),
			"message": err.Error(),
		})
		return
	}
	//dirtytime
	if req.DirtyTimestamp > 0 {
		instance.DirtyTimestamp = req.DirtyTimestamp
	}
	global.Discovery.Registry.Register(instance, req.LatestTimestamp)
	//from other server, Replication is true
	log.Println("replicattion", req.Replication)
	if req.Replication {
		global.Discovery.Nodes.Load().(*model.Nodes).Replicate(c, configs.Register, instance)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "",
		"data":    "",
	})
}
