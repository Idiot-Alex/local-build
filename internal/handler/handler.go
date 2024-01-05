package handler

import (
	"github.com/gin-gonic/gin"
)

// export init env
func InitEnv(c *gin.Context) {
	initEnv(c)
}

// export list tool
func ListTool(c *gin.Context) {
	listTool(c)
}

// export save tool
func SaveTool(c *gin.Context) {
	saveTool(c)
}

// export del tool
func DelTool(c *gin.Context) {
	delTool(c)
}