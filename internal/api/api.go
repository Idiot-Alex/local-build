package api

import (
	"local-build/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitApis(r *gin.Engine) {
	r.GET("/", helloWorld)

	env := r.Group("/api/env")
	{
		env.GET("/init", controller.EnvInit)
	}
	
	tool := r.Group("/api/tool")
  // {} 是书写规范
  {
		tool.GET("/list", controller.ToolList)
		tool.POST("/save", controller.SaveTool)
		tool.GET("/del", controller.DelTool)
  }

	project := r.Group("/api/project")
	{
		tool.GET("/list", controller.ProjectList)
		project.POST("/save", controller.SaveProject)
		project.GET("/del", controller.DelProject)
	}
}


// gin.Context，封装了request和response
func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}

