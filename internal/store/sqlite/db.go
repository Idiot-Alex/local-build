package sqlite

import (
	"local-build/internal/config"
	"local-build/internal/lblog"
	"local-build/internal/store/model"
	"net/url"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// get DB
func GetDB() gorm.DB {
	cfg := config.Load("config.toml")
	dbFilePath := cfg.Db.FilePath
	parsedUrl, err := url.Parse(dbFilePath)
	if err != nil {
		lblog.Error(err)
		panic(err)
	}
	query := parsedUrl.Query()
	query.Set("_json", "1")
	parsedUrl.RawQuery = query.Encode()
	dbFilePath = parsedUrl.String()
	lblog.Info("dbFilePath: %s", dbFilePath)

	db, err := gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		lblog.Error(err)
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		lblog.Error(err)
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
	db.AutoMigrate(&model.Project{}, &model.Tool{}, &model.BuildPlan{}, &model.BuildTask{})
}
