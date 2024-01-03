package controller

import (
	"encoding/json"
	"local-build/internal/model"
	"local-build/internal/sqlite"
	"log"

	"github.com/gin-gonic/gin"
)

// tool list
func ToolList(c *gin.Context) {
	var tools []model.Tool
	db := sqlite.GetDB()
	db.Find(&tools).Offset(0).Limit(10).Order("name")
	
	jsonData, _ := json.MarshalIndent(tools, "", "  ")
	log.Printf("tool list: %+v", string(jsonData))
	c.String(200, string(jsonData))
}

// save tool
func SaveTool(c *gin.Context) {
	
}

// del tool
func DelTool(c *gin.Context) {

}