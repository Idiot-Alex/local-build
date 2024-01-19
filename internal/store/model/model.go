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
	ID         string  `gorm:"primarykey;comment:'ID'" json:"id"`
	Name       string  `gorm:"unique;not null;comment:'项目名称'" json:"name"`
	Desc       string  `gorm:"comment:'项目描述'" json:"desc"`
	Type       string  `gorm:"comment:'项目类型[GIT,Dir]'" json:"type"`
	Config     string  `gorm:"type:json;comment:'项目配置'" json:"config"`
	ParsedInfo string  `gorm:"type:json;comment:'项目解析后的信息'" json:"parsedInfo"`
	CreatedAt  FmtTime `gorm:"comment:'创建时间'" json:"createdAt"`
	UpdatedAt  FmtTime `gorm:"comment:'修改时间'" json:"updatedAt"`
}

type Tool struct {
	ID        string  `gorm:"primarykey;comment:'ID'" json:"id"`
	Name      string  `gorm:"unique;not null;comment:'工具名称'" json:"name" binding:"required"`
	Desc      string  `gorm:"comment:'工具描述'" json:"desc"`
	Path      string  `gorm:"comment:'工具路径'" json:"path" binding:"required"`
	Version   string  `gorm:"version;not null;comment:'版本号'" json:"version"`
	Vendor    string  `gorm:"comment:'厂商'" json:"vendor"`
	Arch      string  `gorm:"comment:'系统架构'" json:"arch"`
	Type      string  `gorm:"comment:'类型[GIT,MAVEN,JDK,NODE]'" json:"type" binding:"required"`
	Config    string  `gorm:"type:json;comment:'配置信息'" json:"config"`
	CreatedAt FmtTime `gorm:"comment:'创建时间'" json:"createdAt"`
	UpdatedAt FmtTime `gorm:"comment:'修改时间'" json:"updatedAt"`
}

type BuildPlan struct {
	ID         string  `gorm:"primarykey;comment:'ID'" json:"id"`
	Name       string  `gorm:"unique;not null;comment:'构建计划名称'" json:"name" binding:"required"`
	PId        string  `gorm:"comment:'父级ID'" json:"pId"`
	LastCostMs int64   `gorm:"comment:'上次花费时间毫秒'" json:"lastCostMs"`
	Config     string  `gorm:"type:json;comment:'配置信息'" json:"config"`
	CreatedAt  FmtTime `gorm:"comment:'创建时间'" json:"createdAt"`
	UpdatedAt  FmtTime `gorm:"comment:'修改时间'" json:"updatedAt"`
}

type BuildTask struct {
	ID        string  `gorm:"primarykey;comment:'ID'" json:"id"`
	Num       int64   `gorm:"comment:'构建任务序号'" json:"num"`
	State     string  `gorm:"comment:'状态[SUCCESS,FAIL,ABORT]'" json:"state"`
	CreatedAt FmtTime `gorm:"comment:'创建时间'" json:"createdAt"`
	UpdatedAt FmtTime `gorm:"comment:'修改时间'" json:"updatedAt"`
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
