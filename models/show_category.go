package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type ShowCategory struct {
	gorm.Model
	No          string         `sql:"-" json:"id"`
	TenantId    string         `sql:"type:char(20);index" description:"租户ID" json:"-"`
	Name        string         `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Image       string         `json:"image"`
	Parent      *ShowCategory  `gorm:"save_associations:false" json:"parent"`
	ParentId    string         `sql:"type:char(20);index" json:"parent_id" description:"父级ID"`
	Status      Status         `sql:"type:integer;default(1);index" description:"展示状态" json:"status" validate:"required"`
	Description string         `sql:"type:text" description:"描述" json:"description"`
	Children    []ShowCategory `gorm:"ForeignKey:ParentId" json:"children"`
	Sort        int            `description:"排序" json:"sort"`
}

//只用于查找
type ShowCategoryParentGoods struct {
	gorm.Model
	No          string  `sql:"-" json:"id"`
	TenantId    string  `sql:"type:char(20);index" description:"租户ID" json:"-"`
	Name        string  `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Image       string  `json:"image"`
	Status      Status  `sql:"type:integer;default(1);index" description:"展示状态" json:"status" validate:"required"`
	Description string  `sql:"type:text" description:"描述" json:"description"`
	Goods       []Goods `gorm:"ForeignKey:ParentShowCategoryId" json:"goods"`
	Sort        int     `description:"排序" json:"sort"`
}

func (ShowCategoryParentGoods) TableName() string {
	return "show_categories"
}

type ShowCategoryGoods struct {
	gorm.Model
	No          string  `sql:"-" json:"id"`
	TenantId    string  `sql:"type:char(20);index" description:"租户ID" json:"-"`
	Name        string  `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Image       string  `json:"image"`
	ParentId    string  `sql:"type:char(20);index" json:"parent_id" description:"父级ID"`
	Status      Status  `sql:"type:integer;default(1);index" description:"展示状态" json:"status" validate:"required"`
	Description string  `sql:"type:text" description:"描述" json:"description"`
	Sort        int     `description:"排序" json:"sort"`
	Goods       []Goods `gorm:"ForeignKey:ShowCategoryId" json:"goods"`
}

func (ShowCategoryGoods) TableName() string {
	return "show_categories"
}

func (s *ShowCategory) AfterSave() error {
	s.transform()
	return nil
}

func (s *ShowCategory) AfterFind() error {
	s.transform()
	return nil
}

func (s *ShowCategory) transform() {
	s.No = strconv.Itoa(int(s.ID))
}

func (s *ShowCategoryParentGoods) transform() {
	s.No = strconv.Itoa(int(s.ID))
}

func (s *ShowCategoryParentGoods) AfterSave() error {
	s.transform()
	return nil
}

func (s *ShowCategoryParentGoods) AfterFind() error {
	s.transform()
	return nil
}

type SearchShowCategory struct {
	List      []ShowCategory `json:"list"`
	Total     int            `json:"total"`
	Page      int            `json:"page"`
	TotalPage int            `json:"total_page"`
	Limit     int            `json:"limit"`
}
