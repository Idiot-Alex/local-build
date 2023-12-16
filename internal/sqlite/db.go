package sqlite

import (
	"database/sql"
	"log"
	"reflect"
	"runtime"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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

