package service

import (
	"fmt"
	"local-build/internal/fileparse"
	"local-build/internal/lblog"
	"local-build/internal/pkg/env"
	"local-build/internal/repo"
	"local-build/internal/store/model"
	"local-build/internal/store/sqlite"
	"local-build/internal/utils"
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
			lblog.Error("%s records found...name: %s, path: %s", v.Type, v.Name, v.Path)
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
			lblog.Error("%s records found...name: %s, path: %s", tool.Type, tool.Name, tool.Path)
			return false
		}
		tool.ID = utils.GenerateIdStr()
	}

	lblog.Info("tool: %s", utils.ToJsonString(tool))
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
	// parse project
	parsedInfo := ParseProject(project)
	project.ParsedInfo = utils.ToJsonString(parsedInfo)

	db := sqlite.GetDB()

	// check name before insert
	if project.ID == "" {
		var count int64
		err := db.Model(&project).Where("name = ?", project.Name).Count(&count).Error
		if err != nil {
			panic(err)
		}

		if count != 0 {
			lblog.Error("project records found...name: %s, path: %s", project.Name, project.Path)
			return false
		}
		project.ID = utils.GenerateIdStr()
	}

	lblog.Info("project: %s", utils.ToJsonString(project))
	tx := db.Save(&project)
	return tx.RowsAffected > 0
}

// parse project
func ParseProject(p model.Project) *fileparse.ParsedInfo {
	// assert project repo type
	switch p.RepoConfig.Type {
	case env.GIT:
		conf := repo.GitConfig{
			Url:           p.RepoConfig.Url,
			LocalPath:     p.Path,
			AccessType:    p.RepoConfig.AccessType,
			UserName:      p.RepoConfig.UserName,
			Password:      p.RepoConfig.Password,
			SshPrivateKey: p.RepoConfig.SshPrivateKey,
			KeyPassphrase: p.RepoConfig.KeyPassphrase,
			AccessToken:   p.RepoConfig.AccessToken,
		}

		err := repo.UpdateGitRepo(conf)
		if err != nil {
			lblog.Error("update git repo error: %s", err)
			panic(err)
		}
	case env.DIR:
		// check if path exist
		dir := filepath.Join(p.Path)
		_, err := os.Stat(dir)
		if err != nil {
			lblog.Error("the file path [%s] error: %s ", dir, err)
			panic(err)
		}
	case env.SVN:
	}

	// parse directory
	parsed, err := fileparse.ParseDirectory(p.Path)
	if err != nil {
		lblog.Error("parse directory [%s] error: %s", p.Path, err)
		panic(err)
	}
	lblog.Info("parsed info: %s", utils.ToJsonString(parsed))

	return parsed
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

	// check name and projectId before insert
	if buildPlan.ID == "" {
		var count int64
		err := db.Model(&buildPlan).Where("name = ? and project_id = ?", buildPlan.Name, buildPlan.ProjectId).Count(&count).Error
		if err != nil {
			panic(err)
		}

		if count != 0 {
			lblog.Error("build plan records found...name: %s, projectId: %s", buildPlan.Name, buildPlan.ProjectId)
			return false
		}
		buildPlan.ID = utils.GenerateIdStr()
	}

	lblog.Info("build plan: %s", utils.ToJsonString(buildPlan))
	tx := db.Save(&buildPlan)
	return tx.RowsAffected > 0
}

// del build plan
func DelBuildPlan(ids []string) bool {
	db := sqlite.GetDB()
	tx := db.Where("id in ?", ids).Delete(&model.BuildPlan{})
	return tx.RowsAffected > 0
}
