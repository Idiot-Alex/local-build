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
		env.GET("/init", handler.EnvInit)
	}
	
	tool := r.Group("/api/tool")
  // {} 是书写规范
  {
		tool.GET("/list", handler.ToolList)
		tool.POST("/save", handler.SaveTool)
		tool.GET("/del", handler.DelTool)
  }

	project := r.Group("/api/project")
	{
		tool.GET("/list", handler.ProjectList)
		project.POST("/save", handler.SaveProject)
		project.GET("/del", handler.DelProject)
	}
}


// gin.Context，封装了request和response
func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}

