package service

import (
	"local-build/internal/pkg/env"
	"local-build/internal/store/model"
	"local-build/internal/store/sqlite"
	"log"
)

// export init tools
func InitTools() {
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

// export tool list
func ToolList() []model.Tool {
	var tools []model.Tool
	db := sqlite.GetDB()
	db.Find(&tools).Offset(0).Limit(10).Order("name")
	return tools
}

// export save tool
func SaveTool() bool {
	return true
}

// export del tool
func DelTool(ids []string) bool {
	log.Printf("ids: %v\n", ids)
	db := sqlite.GetDB()
	tx := db.Where("id in ?", ids).Delete(&model.Tool{})
	return tx.RowsAffected > 0
}