package api

import (
	"local-build/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitApis(r *gin.Engine) {
	r.GET("/", helloWorld)

	env := r.Group("/api/env")
	{
		env.GET("/init", handler.InitEnv)
	}

	tool := r.Group("/api/tool")
	// {} 是书写规范
	{
		tool.POST("/list", handler.ToolList)
		tool.POST("/save", handler.SaveTool)
		tool.POST("/del", handler.DelTool)
	}

	project := r.Group("/api/project")
	{
		project.POST("/list", handler.ProjectList)
		project.POST("/save", handler.SaveProject)
		project.POST("/del", handler.DelProject)
	}

	plan := r.Group("/api/build-plan")
	{
		plan.POST("/list", handler.BuildPlanList)
		plan.POST("/save", handler.SaveBuildPlan)
		plan.POST("/del", handler.DelBuildPlan)
	}

	task := r.Group("/api/build-task")
	{
		task.POST("/list", handler.BuildTaskList)
		task.POST("/save", handler.SaveBuildTask)
		task.POST("/del", handler.DelBuildTask)
	}
}

// gin.Context，封装了request和response
func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}
