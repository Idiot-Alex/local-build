package test

import (
	"fmt"
	"local-build/internal/model"
	"local-build/internal/sqlite"
	"local-build/internal/utils"
	"testing"
)

func TestDb(t *testing.T) {
	db := sqlite.GetDB()

	if !db.Migrator().HasTable(&model.Project{}) {
		db.AutoMigrate(&model.Project{})
	}

	p := model.Project{ID: utils.GenerateId(), Name: "xx"}
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

func TestSqlite(t *testing.T) {
	db, err := sqlite.NewDBHelper("local-build.sqlite?_json=1")
	if nil != err {
		t.Errorf(err.Error())
	}

	db.CreateTable("create table if not exists user(id integer primary key, firstname text, lastname text)")

	db.InsertData("insert into user( firstname, lastname) values(?,?)", "s1", "s1")
	db.InsertData("insert into user( firstname, lastname) values(?,?)", "s1", "s2")

	rows, err := db.Query("select id, firstname, lastname from user")
	if nil != err {
			t.Log(err)
	}
	t.Logf("rows: %v\n", rows)
}