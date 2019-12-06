package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

//国际版物流
//20191206兼容国内物流
type Courier struct {
	gorm.Model
	No      string            `sql:"-" json:"id"`
	Name    string            `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Logo    string            `sql:"type:varchar(255)" description:"logo图标" json:"logo"`
	Status  Status            `sql:"default(1)" description:"状态" json:"status"`
	SiteUrl string            `sql:"type:varchar(255)" description:"官网网址" json:"site_url"`
	Region  Region            `sql:"type:varchar(20)" description:"地区" json:"region" validate:"required"`
	Method  string            `sql:"type:varchar(100)" description:"调用方法名" json:"method" validate:"required"`
	Rules   []CourierPackRule `gorm:"ForeignKey:CourierId" json:"rules"`
}

type CourierTemplate struct {
	gorm.Model
	No             string   `sql:"-" json:"id"`
	Courier        Courier  `json:"courier" validate:"-"`
	CourierId      int      `description:"物流ID" json:"courier_id"`
	Name           string   `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	FirstWeight    int      `description:"首重" json:"first_weight"`
	FirstPrice     float64  `sql:"type:DECIMAL(10, 2)" description:"首重价格" json:"first_price"`
	ContinuedPrice float64  `sql:"type:DECIMAL(10, 2)" description:"续重价格" json:"continued_price"`
	CodeData       string   `sql:"type:varchar(100)" description:"区域" json:"code"`
	Code           []string `sql:"-" description:"区域集" json:"code"`
}

type CourierInstall struct {
	gorm.Model
	No        string  `sql:"-" json:"id"`
	Used      bool    `description:"领用" json:"used"`
	TenantId  string  `json:"tenant_id"`
	Courier   Courier `json:"courier" validate:"-"`
	CourierId int     `description:"物流ID" json:"courier_id"`
	AppKey    string  `sql:"type:varchar(50)" description:"key" json:"app_key"`
	AppSecret string  `sql:"type:varchar(50)" description:"密钥" json:"app_secret"`
	MaxAmount float64 `sql:"type:DECIMAL(10, 2)" description:"最大打包金额" json:"max_amount"`
	MaxWeight int     `description:"包裹最大重量" json:"max_weight"`
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
	Courier    Courier `validate:"-" json:"courier"`
	CourierId  string  `json:"courier_id"`
	No         string  `sql:"-" json:"id"`
	Name       string  `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Simple     int     `sql:"default(0)" description:"纯装数量" json:"simple"`
	Mixed      int     `sql:"default(0)" description:"混装基数" json:"mixed"`
	MixedSum   int     `sql:"default(0)" description:"混装总数" json:"mixed_sum"`
	PriceUnit  float64 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"单价" json:"price_unit"`
	PriceTotal float64 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"总价" json:"price_total"`
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
	return nil
}

func (c *CourierPackRule) AfterSave() error {
	c.No = strconv.Itoa(int(c.ID))
	return nil
}

func (c *CourierInstall) AfterFind() error {
	c.No = strconv.Itoa(int(c.ID))
	return nil
}

func (c *CourierInstall) AfterSave() error {
	c.No = strconv.Itoa(int(c.ID))
	return nil
}
