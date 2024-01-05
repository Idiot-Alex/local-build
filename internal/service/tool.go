package service

import (
	"local-build/internal/env"
	"local-build/internal/store/model"
	"local-build/internal/store/sqlite"
	"log"
)

// init tools
func initTools() {
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

	node := env.GetNodeInfo()
	if node != (model.Tool{}) {
		tools = append(tools, node)
	}

	db := sqlite.GetDB()
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

// tool list
func toolList() []model.Tool {
	var tools []model.Tool
	db := sqlite.GetDB()
	db.Find(&tools).Offset(0).Limit(10).Order("name")
	return tools
}

// save tool
func saveTool() {

}

// del tool
func delTool(ids []string) bool {
	db := sqlite.GetDB()
	db.Delete(&model.Tool{}, ids)
	return db.RowsAffected > 0
}