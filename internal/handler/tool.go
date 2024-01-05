package handler

import (
	"local-build/internal/service"
	"local-build/internal/store/model"
	"log"

	"github.com/gin-gonic/gin"
)

// list tool
func listTool(c *gin.Context) {
	tools := service.ToolList()
	res := model.Res{Msg: "success", Data: tools}
	log.Printf("tool list: %+v\n", res)
	c.JSON(200, res)
}

// save tool
func saveTool(c *gin.Context) {
	
}

// del tool
func delTool(c *gin.Context) {

}