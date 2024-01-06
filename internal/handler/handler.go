package handler

import (
	"local-build/internal/service"
	"local-build/internal/store/model"
	"log"

	"github.com/gin-gonic/gin"
)

// export init env
func InitEnv(c *gin.Context) {
	service.InitTools()
	res := model.Res{Msg: "init success"}
	log.Printf("env init...res: %+v\n", res)
	c.JSON(200, res)
}

// export list tool
func ListTool(c *gin.Context) {
	tools := service.ToolList()
	res := model.Res{Msg: "success", Data: tools}
	log.Printf("tool list: %+v\n", res)
	c.JSON(200, res)
}

// export save tool
func SaveTool(c *gin.Context) {
	
}

// export del tool
func DelTool(c *gin.Context) {
	var req model.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	var res model.Res
	if service.DelTool(req.Ids) {
		res = model.Res{Msg: "del tool success"}
	} else {
		res = model.Res{Msg: "del tool failed"}
	}
	log.Printf("del tool: %+v\n", res)
	c.JSON(200, res)
}