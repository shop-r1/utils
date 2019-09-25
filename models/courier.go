package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

//国际版物流
type Courier struct {
	gorm.Model
	No      string            `sql:"-" json:"id"`
	Name    string            `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Status  Status            `sql:"default(1)" description:"状态" json:"status" validate:"required"`
	SiteUrl string            `sql:"type:varchar(255)" description:"官网网址" json:"site_url"`
	Rules   []CourierPackRule `gorm:"ForeignKey:CourierId" json:"rules"`
}

type SearchCourier struct {
	List      []Courier `json:"list"`
	Total     int       `json:"total"`
	Page      int       `json:"page"`
	TotalPage int       `json:"total_page"`
	Limit     int       `json:"limit"`
}

type CourierPackRule struct {
	gorm.Model
	Courier       Courier `validate:"-" json:"courier"`
	CourierId     uint    `json:"courier_id_int"`
	CourierNo     string  `sql:"-" json:"courier_id"`
	No            string  `sql:"-" json:"id"`
	Name          string  `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Simple        int     `sql:"default(0)" description:"纯装数量" json:"simple"`
	Mixed         int     `sql:"default(0)" description:"混装基数" json:"mixed"`
	MixedCan      int     `sql:"default(0)" description:"可混装数" json:"mixed_can"`
	MixedSum      int     `sql:"default(0)" description:"混装总数" json:"mixed_sum"`
	PriceUnitAud  float64 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"单价(澳币)" json:"price_unit_aud"`
	PriceUnitRmb  float64 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"单价(人民币)" json:"price_unit_rmb"`
	PriceTotalAud float64 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"总价(澳币)" json:"price_total_aud"`
	PriceTotalRmb float64 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"总价(人民币)" json:"price_total_rmb"`
}

type SearchCourierPackRule struct {
	List      []CourierPackRule `json:"list"`
	Total     int               `json:"total"`
	Page      int               `json:"page"`
	TotalPage int               `json:"total_page"`
	Limit     int               `json:"limit"`
}

func (c *Courier) AfterFind() error {
	c.No = strconv.Itoa(int(c.ID))
	return nil
}

func (c *Courier) AfterSave() error {
	c.No = strconv.Itoa(int(c.ID))
	return nil
}

func (c *CourierPackRule) AfterFind() error {
	c.No = strconv.Itoa(int(c.ID))
	c.CourierNo = strconv.Itoa(int(c.CourierId))
	return nil
}

func (c *CourierPackRule) AfterSave() error {
	c.No = strconv.Itoa(int(c.ID))
	return nil
}
