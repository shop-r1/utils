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
	ParentId    string         `sql:"type:char(20);index" json:"parent_id" description:"父级ID"`
	Status      Status         `sql:"type:integer;default(1);index" description:"展示状态" json:"status" validate:"required"`
	Description string         `sql:"type:text" description:"描述" json:"description"`
	Children    []ShowCategory `gorm:"ForeignKey:ParentId" json:"children"`
	Sort        int            `description:"排序" json:"sort"`
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

type SearchShowCategory struct {
	List      []ShowCategory `json:"list"`
	Total     int            `json:"total"`
	Page      int            `json:"page"`
	TotalPage int            `json:"total_page"`
	Limit     int            `json:"limit"`
}
