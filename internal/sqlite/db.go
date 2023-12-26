package sqlite

import (
	"local-build/internal/model"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetDB() gorm.DB {
	db, err := gorm.Open(sqlite.Open("local-build.sqlite?_json=1"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
  if err != nil {
    panic("failed to connect database")
  }
	sqlDB, err := db.DB()
	if err != nil {
    panic("failed to connect database")
  }
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return *db
}

func InitTable() {
	db := GetDB()
	db.AutoMigrate(&model.Project{}, &model.Tool{})
}
