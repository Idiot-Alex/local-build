package test

import (
	"fmt"
	"local-build/internal/store/model"
	"local-build/internal/store/sqlite"
	"local-build/internal/utils"
	"testing"
)

func TestDb(t *testing.T) {
	// cfg := config.Load("resources/test_config.toml")
	db := sqlite.GetDB()

	sqlite.InitTable()

	p := model.Project{ID: utils.GenerateIdStr(), Name: "xx"}
	res := db.Create(&p)

	fmt.Printf("error: %v\n", res.Error)
	fmt.Printf("id: %v\n", p.ID)
	fmt.Printf("affected: %v\n", res.RowsAffected)

	var project model.Project 
	db.First(&project)

	fmt.Printf("%#v\n", project)

	var count int64
	db.First(&project).Count(&count)

	fmt.Printf("%#v\n", count)
}
