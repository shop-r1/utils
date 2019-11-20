package models

import (
	"encoding/json"
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

type WarehouseCourier struct {
	Id        string `json:"id"`
	CourierId string `json:"courier_id"`
	Name      string `json:"name"`
	Default   bool   `json:"default"`
}

//发货仓
type ShippingWarehouse struct {
	gorm.Model
	TenantId     string             `sql:"type:char(20);index" json:"-"`
	No           string             `sql:"-" json:"id"`
	Name         string             `sql:"type:varchar(100)" description:"发货仓名称" json:"name" validate:"required"`
	Currency     Currency           `sql:"type:varchar(20)" description:"币种" json:"currency" validate:"required"`
	Region       Region             `sql:"type:varchar(20)" description:"地区" json:"region" validate:"required"`
	Address      string             `sql:"type:text" description:"真实地址" json:"address" validate:"required"`
	Status       Status             `sql:"type:integer;default(1)" description:"状态 1启用 2禁用" json:"status"`
	GetSelf      bool               `description:"自提" json:"get_self"`
	NeedIdCard   bool               `description:"需要身份证照片必传" json:"need_id_card"`
	Couriers     []WarehouseCourier `sql:"-" description:"关联物流" json:"couriers"`
	CouriersData []byte             `sql:"type:json" json:"-"`
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

func (s *ShippingWarehouse) BeforeSave() error {
	s.unTransform()
	return nil
}

func (s *ShippingWarehouse) transform() {
	s.No = strconv.Itoa(int(s.ID))
	_ = json.Unmarshal(s.CouriersData, &s.Couriers)
}

func (s *ShippingWarehouse) unTransform() {
	if s.ID == 0 && s.No != "" && s.No != "0" {
		id, _ := strconv.Atoi(s.No)
		s.ID = uint(id)
	}
	if len(s.Couriers) > 0 {
		s.CouriersData, _ = json.Marshal(s.Couriers)
	} else {
		s.CouriersData = []byte(`[]`)
	}
}
