package api

import (
	"local-build/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitApis(r *gin.Engine) {
	r.GET("/", helloWorld)
	
	env := r.Group("/api/env")
  // {} 是书写规范
  {
		env.POST("/list", controller.EnvList)
  }
}


// gin.Context，封装了request和response
func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}

