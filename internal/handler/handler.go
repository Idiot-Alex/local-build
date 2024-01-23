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
	var toolQuery model.ToolQuery
	if err := c.ShouldBindJSON(&toolQuery); err != nil {
		c.JSON(400, &model.Res{Msg: err.Error()})
		return
	}

	tools := service.ToolList(&toolQuery)
	res := model.Res{Msg: "success", Data: tools}
	log.Printf("tool list: %#v\n", res)
	c.JSON(200, res)
}

// export save tool
func SaveTool(c *gin.Context) {
	var tool model.Tool
	if err := c.ShouldBindJSON(&tool); err != nil {
		c.JSON(400, &model.Res{Msg: err.Error()})
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
		c.JSON(400, &model.Res{Msg: err.Error()})
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
	var projectQuery model.ProjectQuery
	if err := c.ShouldBindJSON(&projectQuery); err != nil {
		c.JSON(400, &model.Res{Msg: err.Error()})
		return
	}

	projects := service.ProjectList(&projectQuery)
	res := model.Res{Msg: "success", Data: projects}
	log.Printf("project list: %+v\n", res)
	c.JSON(200, res)
}

// save project
func SaveProject(c *gin.Context) {
	var project model.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(400, &model.Res{Msg: err.Error()})
		return
	}

	var res model.Res
	if service.SaveProject(project) {
		res = model.Res{Msg: "save project success"}
	} else {
		res = model.Res{Msg: "save project failed"}
	}
	log.Printf("save project: %+v\n", res)
	c.JSON(200, res)
}

// del project
func DelProject(c *gin.Context) {
	var req model.ReqIds
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, &model.Res{Msg: err.Error()})
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

// build plan list
func BuildPlanList(c *gin.Context) {
	var query model.BuildPlanQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(400, &model.Res{Msg: err.Error()})
		return
	}

	projects := service.BuildPlanList(&query)
	res := model.Res{Msg: "success", Data: projects}
	log.Printf("build plan list: %+v\n", res)
	c.JSON(200, res)
}

// save build plan
func SaveBuildPlan(c *gin.Context) {
	var buildPlan model.BuildPlan
	if err := c.ShouldBindJSON(&buildPlan); err != nil {
		c.JSON(400, &model.Res{Msg: err.Error()})
		return
	}

	var res model.Res
	if service.SaveBuildPlan(buildPlan) {
		res = model.Res{Msg: "save build plan success"}
	} else {
		res = model.Res{Msg: "save build plan failed"}
	}
	log.Printf("save build plan: %+v\n", res)
	c.JSON(200, res)
}

// del build plan
func DelBuildPlan(c *gin.Context) {
	var req model.ReqIds
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, &model.Res{Msg: err.Error()})
		return
	}

	var res model.Res
	if service.DelBuildPlan(req.Ids) {
		res = model.Res{Msg: "del build plan success"}
	} else {
		res = model.Res{Msg: "del build plan failed"}
	}
	log.Printf("del project: %+v\n", res)
	c.JSON(200, res)
}
