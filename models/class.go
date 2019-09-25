package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Class struct {
	gorm.Model
	Category        Category    `json:"category" validate:"-"`
	No              string      `sql:"-" json:"id"`
	CategoryId      uint        `json:"category_id_int"`
	CategoryNo      string      `sql:"-" json:"category_id"`
	Name            string      `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Attributes      []byte      `sql:"jsonb" description:"属性分组" json:"-"`
	AttributesArray []Attribute `sql:"-" json:"attributes"`
	Status          Status      `sql:"default(1)" description:"状态" json:"status"`
}

type Attribute struct {
	Name       string     `json:"name"`
	Items      []string   `json:"items"`
	SelectType SelectType `json:"select_type"`
}

type SelectType string

const (
	Radio    SelectType = "radio"
	Multiple SelectType = "multiple"
)

type SearchClass struct {
	List      []Class `json:"list"`
	Total     int     `json:"total"`
	Page      int     `json:"page"`
	TotalPage int     `json:"total_page"`
	Limit     int     `json:"limit"`
}

func (c *Class) AfterFind() error {
	c.No = strconv.Itoa(int(c.ID))
	c.CategoryNo = strconv.Itoa(int(c.CategoryId))
	_ = json.Unmarshal(c.Attributes, &c.AttributesArray)
	return nil
}

func (c *Class) AfterSave() error {
	c.No = strconv.Itoa(int(c.ID))
	c.CategoryNo = strconv.Itoa(int(c.CategoryId))
	return nil
}

func (c *Class) BeforeSave() error {
	if c.No != "" && c.ID == 0 {
		var id int
		id, _ = strconv.Atoi(c.No)
		c.ID = uint(id)
	}
	if len(c.AttributesArray) > 0 {
		c.Attributes, _ = json.Marshal(c.AttributesArray)
	}
	return nil
}
