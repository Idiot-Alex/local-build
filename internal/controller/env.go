package controller

import (
	"fmt"
	"local-build/internal/env"
	"local-build/internal/model"
	"local-build/internal/sqlite"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// init env
func EnvInit(c *gin.Context) {
	db := sqlite.GetDB()

	// init tools
	initTools(&db)
}

// env list 
// such as: jdk maven node...
func EnvList(c *gin.Context) {
	c.String(200, fmt.Sprintf("hello %s\n", ".."))
}

// init tool
func initTools(db *gorm.DB) {
	var count int64
	tools := env.GetJDKList()

	git := env.GetGitInfo()
	if git != (model.Tool{}) {
		tools = append(tools, git)
	}
	maven := env.GetMavenInfo()
	if maven != (model.Tool{}) {
		tools = append(tools, maven)
	}

	node := env.GetMavenInfo()
	if node != (model.Tool{}) {
		tools = append(tools, node)
	}

	for _, v := range tools {
		err := db.Model(&v).Where("name = ? and version = ?", v.Name, v.Version).Count(&count).Error
		if err != nil {
			panic(err)
		}

		if count != 0 {
			log.Printf("%s records found...name: %s, version: %s", v.Type, v.Name, v.Version)
		} else {
			db.Create(&v)
		}
	}
}