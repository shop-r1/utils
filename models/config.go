package models

import "github.com/jinzhu/gorm"

type SystemConfig struct {
	gorm.Model
	Name     string                 `sql:"type:varchar(100);index" description:"名称" json:"name" validate:"required"`
	Meta     map[string]interface{} `sql:"-" description:"数据集合" json:"meta"`
	Metadata []byte                 `sql:"type:json" description:"数据集合" json:"-"`
}
