package service

import (
	"fmt"
	"local-build/internal/lblog"
	"local-build/internal/pkg/env"
	"local-build/internal/store/model"
	"local-build/internal/store/sqlite"
	"local-build/internal/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
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
			v.ID = utils.GenerateIdStr()
			db.Save(&v)
		}
	}
}

// export tool list
func ToolList(query *model.ToolQuery) *model.PaginationRes {
	offset := (query.PageNo - 1) * query.PageSize
	var tools []model.Tool
	var count int64

	var conditions []string
	if query.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name LIKE '%%%s%%'", query.Name))
	}
	whereSql := strings.Join(conditions, " AND ")

	db := sqlite.GetDB()
	listQuery := db.Model(&model.Tool{}).Offset(offset).Limit(query.PageSize).Order("name")
	listQuery.Where(whereSql).Find(&tools)

	db.Model(&model.Tool{}).Where(whereSql).Count(&count)

	return &model.PaginationRes{DataList: &tools, Total: count}
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

	db := sqlite.GetDB()

	// check name before insert
	if tool.ID == "" {
		var count int64
		err := db.Model(&tool).Where("path = ?", tool.Path).Count(&count).Error
		if err != nil {
			panic(err)
		}

		if count != 0 {
			log.Printf("%s records found...name: %s, path: %s", tool.Type, tool.Name, tool.Path)
			return false
		}
		tool.ID = utils.GenerateIdStr()
	}

	log.Printf("tool: %+v", tool)
	tx := db.Save(&tool)
	return tx.RowsAffected > 0
}

// export del tool
func DelTool(ids []string) bool {
	db := sqlite.GetDB()
	tx := db.Where("id in ?", ids).Delete(&model.Tool{})
	return tx.RowsAffected > 0
}

// export project list
func ProjectList(query *model.ProjectQuery) *model.PaginationRes {
	offset := (query.PageNo - 1) * query.PageSize
	var projects []model.Project
	var count int64

	var conditions []string
	if query.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name LIKE '%%%s%%'", query.Name))
	}
	whereSql := strings.Join(conditions, " AND ")

	db := sqlite.GetDB()

	listQuery := db.Model(&model.Project{}).Offset(offset).Limit(query.PageSize).Order("name")
	listQuery.Where(whereSql).Find(&projects)

	db.Model(&model.Project{}).Where(whereSql).Count(&count)

	return &model.PaginationRes{DataList: &projects, Total: count}
}

// export del tool
func DelProject(ids []string) bool {
	db := sqlite.GetDB()
	tx := db.Where("id in ?", ids).Delete(&model.Project{})
	return tx.RowsAffected > 0
}

// export save project
func SaveProject(project model.Project) bool {
	db := sqlite.GetDB()

	// check name before insert
	if project.ID == "" {
		var count int64
		err := db.Model(&project).Where("name = ?", project.Name).Count(&count).Error
		if err != nil {
			panic(err)
		}

		if count != 0 {
			log.Printf("project records found...name: %s, path: %s", project.Name, project.Path)
			return false
		}
		project.ID = utils.GenerateIdStr()
	}

	log.Printf("project: %+v", project)
	tx := db.Save(&project)
	return tx.RowsAffected > 0
}

// parse project
func ParseProject(p model.Project) string {
	// assert project repoType
	switch p.RepoConfig.AccessType {
	case env.GIT:
		// check if .git exist
		dir := filepath.Join(p.Path, ".git")
		_, err := os.Stat(dir)
		if err != nil {
			if os.IsNotExist(err) {
				lblog.Warning("Directory %s does not exist.", dir)
				// exec git clone
				conf := utils.GitConfig{
					Url:           p.RepoConfig.Url,
					LocalPath:     p.Path,
					AccessType:    p.RepoConfig.AccessType,
					UserName:      p.RepoConfig.UserName,
					Password:      p.RepoConfig.Password,
					SshPrivateKey: p.RepoConfig.SshPrivateKey,
					KeyPassphrase: p.RepoConfig.KeyPassphrase,
					AccessToken:   p.RepoConfig.AccessToken,
				}
				err = utils.GitClone(conf)
				if err != nil {
					lblog.Error("git clone [%s] error: %s", p.Path, err)
					panic(err)
				}
			} else {
				lblog.Error("the file path [%s] error: %s ", dir, err)
				panic(err)
			}
		}

		lblog.Info("Directory %s exists.", dir)
		// foreach dir tree to parse files
		walkFunc := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err // 如果出错，直接返回错误
			}
			if !info.IsDir() {
				lblog.Info("dir: [%s]", path) // 如果是目录，打印目录名
			} else {
				lblog.Info("file: [%s]", path) // 如果是文件，打印文件名
			}
			return nil
		}

		// 遍历目录（包括子目录），调用回调函数处理每个文件或目录
		if err := filepath.Walk(p.Path, walkFunc); err != nil {
			lblog.Error("Walk() failed: %v\n", err)
		}
	case env.DIR:
	case env.SVN:
	}
	return ""
}

// build plan list
func BuildPlanList(query *model.BuildPlanQuery) *model.PaginationRes {
	offset := (query.PageNo - 1) * query.PageSize
	var buildPlans []model.BuildPlan
	var count int64

	var conditions []string
	if query.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name LIKE '%%%s%%'", query.Name))
	}
	whereSql := strings.Join(conditions, " AND ")

	db := sqlite.GetDB()

	listQuery := db.Model(&model.BuildPlan{}).Offset(offset).Limit(query.PageSize).Order("name")
	listQuery.Where(whereSql).Find(&buildPlans)

	db.Model(&model.BuildPlan{}).Where(whereSql).Count(&count)

	return &model.PaginationRes{DataList: &buildPlans, Total: count}
}

// save build plan
func SaveBuildPlan(buildPlan model.BuildPlan) bool {
	db := sqlite.GetDB()

	// check name before insert
	if buildPlan.ID == "" {
		var count int64
		err := db.Model(&buildPlan).Where("name = ?", buildPlan.Name).Count(&count).Error
		if err != nil {
			panic(err)
		}

		if count != 0 {
			log.Printf("build plan records found...name: %s", buildPlan.Name)
			return false
		}
		buildPlan.ID = utils.GenerateIdStr()
	}

	log.Printf("build plan: %+v", buildPlan)
	tx := db.Save(&buildPlan)
	return tx.RowsAffected > 0
}

// del build plan
func DelBuildPlan(ids []string) bool {
	db := sqlite.GetDB()
	tx := db.Where("id in ?", ids).Delete(&model.BuildPlan{})
	return tx.RowsAffected > 0
}
