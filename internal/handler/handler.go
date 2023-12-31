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
func ToolList(c *gin.Context) {
	tools := service.ToolList()
	res := model.Res{Msg: "success", Data: tools}
	log.Printf("tool list: %+v\n", res)
	c.JSON(200, res)
}

// export save tool
func SaveTool(c *gin.Context) {
	var tool model.Tool
	if err := c.ShouldBindJSON(&tool); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var res model.Res
	if service.SaveTool(tool) {
		res = model.Res{Msg: "save tool success"}
	} else {
		res = model.Res{Msg: "save tool failed"}
	}
	log.Printf("save tool: %+v\n", res)
	c.JSON(200, res)
}

// export del tool
func DelTool(c *gin.Context) {
	var req model.ReqIds
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

// project list
func ProjectList(c *gin.Context) {
	projects := service.ProjectList()
	res := model.Res{Msg: "success", Data: projects}
	log.Printf("project list: %+v\n", res)
	c.JSON(200, res)
}

// save project
func SaveProject(c *gin.Context) {

}

// del project
func DelProject(c *gin.Context) {
	var req model.ReqIds
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var res model.Res
	if service.DelProject(req.Ids) {
		res = model.Res{Msg: "del project success"}
	} else {
		res = model.Res{Msg: "del project failed"}
	}
	log.Printf("del project: %+v\n", res)
	c.JSON(200, res)
}
