package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type SystemConfig struct {
	gorm.Model
	TenantId string                 `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	Name     string                 `sql:"type:varchar(100);index" description:"名称" json:"name" validate:"required"`
	Meta     map[string]interface{} `sql:"-" description:"数据集合" json:"meta"`
	Metadata []byte                 `sql:"type:json" description:"数据集合" json:"-"`
}

type SearchSystemConfig struct {
	List      []SystemConfig `json:"list"`
	Total     int            `json:"total"`
	Page      int            `json:"page"`
	TotalPage int            `json:"total_page"`
	Limit     int            `json:"limit"`
}

func (e *SystemConfig) BeforeSave() (err error) {
	if len(e.Meta) == 0 || e.Meta == nil {
		e.Metadata = []byte(`{}`)
	} else {
		e.Metadata, err = json.Marshal(e.Meta)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *SystemConfig) BeforeFind() (err error) {
	err = json.Unmarshal(e.Metadata, &e.Meta)
	return nil
}
