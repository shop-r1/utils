package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type ShowCategory struct {
	gorm.Model
	No          string         `sql:"-" json:"id"`
	TenantId    uint           `sql:"index" description:"租户ID" json:"-"`
	Name        string         `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Image       string         `json:"image"`
	ParentId    int            `sql:"index" json:"parent_id_int" description:"父级ID"`
	ParentNo    string         `sql:"-" json:"parent_id"`
	Status      Status         `sql:"type:integer;default(1);index" description:"展示状态" json:"status" validate:"required"`
	Description string         `sql:"type:text" description:"描述" json:"description"`
	Children    []ShowCategory `gorm:"ForeignKey:ParentId" json:"children"`
}

func (s *ShowCategory) BeforeSave() (err error) {
	if s.ParentNo != "" && s.ParentNo != "0" {
		s.ParentId, err = strconv.Atoi(s.ParentNo)
	}

	return nil
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
	s.ParentNo = strconv.Itoa(s.ParentId)
	s.No = strconv.Itoa(int(s.ID))
}

type SearchShowCategory struct {
	List      []ShowCategory `json:"list"`
	Total     int            `json:"total"`
	Page      int            `json:"page"`
	TotalPage int            `json:"total_page"`
	Limit     int            `json:"limit"`
}
