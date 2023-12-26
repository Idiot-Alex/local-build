package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// env list 
// such as: jdk maven node...
func EnvList(c *gin.Context) {
	
	c.String(200, fmt.Sprintf("hello %s\n", ".."))
}