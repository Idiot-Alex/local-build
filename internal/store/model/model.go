package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type PaginationReq struct {
	PageNo   int `json:"pageNo" binding:"required,min=1"`
	PageSize int `json:"pageSize" binding:"required,min=1"`
}

type PaginationRes struct {
	DataList interface{} `json:"dataList"`
	Total    int64       `json:"total"`
}

type ToolQuery struct {
	PaginationReq
	Name string `json:"name"`
}

type ProjectQuery struct {
	PaginationReq
	Name string `json:"name"`
}

type ReqIds struct {
	Ids []string `json:"ids" binding:"required"`
}

type Res struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Project struct {
	ID         string  `gorm:"primarykey" json:"id"`
	Name       string  `gorm:"unique;not null" json:"name"`
	Desc       string  `json:"desc"`
	Type       string  `json:"type"`
	Config     string  `gorm:"type:json" json:"config"`
	ParsedInfo string  `gorm:"type:json" json:"parsedInfo"`
	CreatedAt  FmtTime `json:"createdAt"`
	UpdatedAt  FmtTime `json:"updatedAt"`
}

type Tool struct {
	ID        string  `gorm:"primarykey" json:"id"`
	Name      string  `gorm:"unique;not null" json:"name" binding:"required"`
	Desc      string  `json:"desc"`
	Path      string  `json:"path" binding:"required"`
	Version   string  `gorm:"version;not null" json:"version"`
	Vendor    string  `json:"vendor"`
	Arch      string  `json:"arch"`
	Type      string  `json:"type" binding:"required"`
	Config    string  `gorm:"type:json" json:"config"`
	CreatedAt FmtTime `json:"createdAt"`
	UpdatedAt FmtTime `json:"updatedAt"`
}

type BuildPlan struct {
	ID         string  `gorm:"primarykey" json:"id"`
	Name       string  `gorm:"unique;not null" json:"name" binding:"required"`
	PId        string  `json:"pId"`
	LastCostMs int64   `json:"lastCostMs"`
	Config     string  `gorm:"type:json" json:"config"`
	CreatedAt  FmtTime `json:"createdAt"`
	UpdatedAt  FmtTime `json:"updatedAt"`
}

type BuildTask struct {
	ID        string  `gorm:"primarykey" json:"id"`
	Num       int64   `json:"num"`
	State     int8    `json:"state"`
	CreatedAt FmtTime `json:"createdAt"`
	UpdatedAt FmtTime `json:"updatedAt"`
}

type FmtTime time.Time

// overwrite MarshalJSON
func (mt FmtTime) MarshalJSON() ([]byte, error) {
	t := time.Time(mt)
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))), nil
}

// overwrite UnMarshalJSON
func (mt *FmtTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// 将字符串解析为 Time 类型
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}
	*mt = FmtTime(t)
	return nil
}

func (ft FmtTime) Value() (driver.Value, error) {
	if time.Time(ft).IsZero() {
		return nil, nil
	}
	return time.Time(ft).Format("2006/01/02 15:04:05"), nil
}

func (ft *FmtTime) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if ok {
		if t.IsZero() {
			t = time.Now()
		}
		*ft = FmtTime(t)
	}
	return nil
}
