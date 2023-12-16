package test

import (
	"local-build/internal/sqlite"
	"testing"
)

func TestSqlite(t *testing.T) {
	db, err := sqlite.NewDBHelper("local-build.sqlite")
	if nil != err {
		t.Errorf(err.Error())
	}

	db.CreateTable("create table if not exists user(id integer primary key, firstname text, lastname text)")

	db.InsertData("insert into user( firstname, lastname) values(?,?)", "s", "s2")

	rows, err := db.Query("select id, firstname, lastname from user")
	if nil != err {
			t.Log(err)
	}
	t.Logf("rows: %v\n", rows)
}