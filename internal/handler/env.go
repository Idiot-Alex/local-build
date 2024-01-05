package handler

import (
	"fmt"
	"local-build/internal/service"
	"local-build/internal/store/model"
	"log"

	"github.com/gin-gonic/gin"
)

// init env
func EnvInit(c *gin.Context) {
	service.InitTools()
	res := model.Res{Msg: "init success"}
	log.Printf("env init...res: %+v\n", res)
	c.JSON(200, res)
}

// env list 
// such as: jdk maven node...
func EnvList(c *gin.Context) {
	// service.ToolList()
	c.String(200, fmt.Sprintf("hello %s\n", ".."))
}
