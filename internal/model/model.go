package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Project struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	Desc      string    `json:"desc"`
	CreatedAt FmtTime `json:"createAt"`
	UpdatedAt FmtTime `json:"updateAt"`
}

type Tool struct {
	ID        int64  `gorm:"primarykey" json:"id"`
	Name      string `gorm:"uniqueIndex:idx_name_version,uniqueIndex;not null" json:"name"`
	Desc      string `json:"desc"`
	Path      string `json:"path"`
	Version   string `gorm:"uniqueIndex:idx_name_version,uniqueIndex;not null" json:"version"`
	Vendor    string `json:"vendor"`
	Arch      string `json:"arch"`
	Type      string `json:"type"`
	Config    string `gorm:"type:json" json:"config"`
	CreatedAt FmtTime `json:"createAt"`
	UpdatedAt FmtTime `json:"updateAt"`
}

type FmtTime time.Time

// overwrite MarshalJSON
func (mt FmtTime) MarshalJSON() ([]byte, error) {
  t := time.Time(mt)
  return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))), nil
}

func (ft FmtTime) Value() (driver.Value, error) {
  tTime := time.Time(ft)
  return tTime.Format("2006/01/02 15:04:05"), nil
}

func (ft *FmtTime) Scan(value interface{}) error {
  t, ok := value.(time.Time)
  if ok {
    *ft = FmtTime(t)
  }
  return nil
}