package controller

import (
	"fmt"
	"local-build/internal/env"
	"local-build/internal/sqlite"

	"github.com/gin-gonic/gin"
)

// init env
func EnvInit(c *gin.Context) {
	db := sqlite.GetDB()

	// init jdk
	jdkList := env.GetJDKList()
	if len(jdkList) > 0 {
		db.Create(&jdkList)
	}

	// init git

	// init maven

	// init node
}

// env list 
// such as: jdk maven node...
func EnvList(c *gin.Context) {
	c.String(200, fmt.Sprintf("hello %s\n", ".."))
}