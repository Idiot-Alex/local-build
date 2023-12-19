package model

import "time"

type Project struct {
	ID int64 `gorm:"primarykey"`
	Name string `gorm:"not null"`
	Desc string
	CreatedAt time.Time
  UpdatedAt time.Time
}

type Tool struct {
	ID int64 `gorm:"primarykey"`
  Name string `gorm:"not null"`
  Desc string
  Path string
  Version string
  Vendor string
  Arch string
  Config string `gorm:"type:json"`
  CreatedAt time.Time
  UpdatedAt time.Time
}