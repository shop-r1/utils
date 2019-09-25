package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type (
	Currency string
	Region   string
)

const (
	AUD       Currency = "aud"
	RMB       Currency = "rmb"
	China     Region   = "china"
	Australia Region   = "australia"
)

type ShippingWarehouse struct {
	gorm.Model
	TenantId uint     `json:"-"`
	No       string   `sql:"-" json:"id"`
	Name     string   `sql:"type:varchar(100)" json:"name" validate:"required"`
	Type     Currency `sql:"type:varchar(20)" json:"type" validate:"required"`
	Region   Region   `sql:"type:varchar(20)" json:"region" validate:"required"`
	Address  string   `sql:"type:text" json:"address" validate:"required"`
	Status   Status   `sql:"type:integer;default(1)" description:"状态 1启用 2禁用"`
}

type SearchShippingWarehouse struct {
	List      []ShippingWarehouse `json:"list"`
	Total     int                 `json:"total"`
	Page      int                 `json:"page"`
	TotalPage int                 `json:"total_page"`
	Limit     int                 `json:"limit"`
}

func (s *ShippingWarehouse) AfterSave() error {
	s.transform()
	return nil
}

func (s *ShippingWarehouse) AfterFind() error {
	s.transform()
	return nil
}

func (s *ShippingWarehouse) transform() {
	s.No = strconv.Itoa(int(s.ID))
}
