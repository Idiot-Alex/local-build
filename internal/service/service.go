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
		err := db.Model(&v).Where("path = ?", v.Path).Count(&count).Error
		if err != nil {
			panic(err)
		}

		if count != 0 {
			log.Printf("%s records found...name: %s, path: %s", v.Type, v.Name, v.Path)
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
func SaveTool(tool model.Tool) bool {
	var ver string
	switch tool.Type {
	case env.NODE:
		ver = env.NodeVersion(tool.Path)
	case env.MAVEN:
		ver = env.MavenVersion(tool.Path)
	case env.GIT:
		ver = env.GitVersion(tool.Path)
	case env.JDK:
		ver = env.JDKVersion(tool.Path)
	}

	tool.Version = ver
	tool.Arch = env.GetOSArch()

	var count int64
	db := sqlite.GetDB()
	err := db.Model(&tool).Where("path = ?", tool.Path).Count(&count).Error
	if err != nil {
		panic(err)
	}

	if count != 0 {
		log.Printf("%s records found...name: %s, path: %s", tool.Type, tool.Name, tool.Path)
		return false
	}

	tx := db.Create(&tool)
	return tx.RowsAffected > 0
}

// export del tool
func DelTool(ids []string) bool {
	db := sqlite.GetDB()
	tx := db.Where("id in ?", ids).Delete(&model.Tool{})
	return tx.RowsAffected > 0
}

// export project list
func ProjectList() []model.Project {
	var list []model.Project
	db := sqlite.GetDB()
	db.Find(&list).Offset(0).Limit(10).Order("name")
	return list
}

// export del tool
func DelProject(ids []string) bool {
	db := sqlite.GetDB()
	tx := db.Where("id in ?", ids).Delete(&model.Project{})
	return tx.RowsAffected > 0
}