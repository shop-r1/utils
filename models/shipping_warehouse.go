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
	AUD       Currency = "AUD"
	RMB       Currency = "CNY"
	China     Region   = "china"
	Australia Region   = "australia"
)

//发货仓
type ShippingWarehouse struct {
	gorm.Model
	TenantId  string   `sql:"type:char(20);index" json:"-"`
	No        string   `sql:"-" json:"id"`
	Name      string   `sql:"type:varchar(100)" description:"发货仓名称" json:"name" validate:"required"`
	Currency  Currency `sql:"type:varchar(20)" description:"币种" json:"currency" validate:"required"`
	Region    Region   `sql:"type:varchar(20)" description:"地区" json:"region" validate:"required"`
	Address   string   `sql:"type:text" description:"真实地址" json:"address" validate:"required"`
	MaxAmount int      `sql:"type:integer;default(0)" description:"包裹最大金额" json:"max_amount"`
	Status    Status   `sql:"type:integer;default(1)" description:"状态 1启用 2禁用" json:"status"`
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
