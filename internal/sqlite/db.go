package sqlite

import (
	"database/sql"
	"fmt"
	"local-build/internal/model"
	"log"
	"reflect"
	"runtime"
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

	project := &model.Project{}
	if !db.Migrator().HasTable(project) {
		fmt.Println("create table for project")
		db.AutoMigrate(project)
	}

	tool := &model.Tool{}
	if !db.Migrator().HasTable(tool) {
		fmt.Println("create table for tool")
		db.AutoMigrate(tool)
	}
}

type DBHelper struct {
	Db *sql.DB
}

func Timing(f func(*DBHelper, ...interface{}) error) func(*DBHelper, ...interface{}) error {
	return func(d *DBHelper, args ...interface{}) error {
			start := time.Now()
			defer func() {
					log.Printf("%s took %v\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), time.Since(start))
			}()
			return f(d, args...)
	}
}

func NewDBHelper(filePath string) (*DBHelper, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, err
	}
	return &DBHelper{db}, nil
}

func (d *DBHelper) Close() {
	d.Db.Close()
}

func (d *DBHelper) CreateTable(query string) error {
	return Timing(func(d *DBHelper, args ...interface{}) error {
		_, err := d.Db.Exec(query)
		if err != nil {
			return err
		}
		return nil
	})(d, query)
}

func (d *DBHelper) InsertData(query string, args ...interface{}) error {
	return Timing(func(d *DBHelper, args ...interface{}) error {
		stmt, err := d.Db.Prepare(query)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
		return nil
	})(d, args...)
}

func (d *DBHelper) UpdateData(query string, args ...interface{}) error {
	return Timing(func(d *DBHelper, args ...interface{}) error {
		stmt, err := d.Db.Prepare(query)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
		return nil
	})(d, args...)
}

func (d *DBHelper) DeleteData(query string, args ...interface{}) error {
	return Timing(func(d *DBHelper, args ...interface{}) error {
		stmt, err := d.Db.Prepare(query)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
		return nil
	})(d, args...)
}

func (d *DBHelper) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
		rows, err := d.Db.Query(query, args...)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}

		var result []map[string]interface{}
		values := make([]interface{}, len(columns))

		for rows.Next() {
			for i := range columns {
				values[i] = new(interface{})
			}
			err := rows.Scan(values...)
			if err != nil {
				return nil, err
			}
			row := make(map[string]interface{})
			for i, col := range columns {
				val := values[i].(*interface{})
				row[col] = *val
			}
			result = append(result, row)
		}

		return result, nil
}

