package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type Payment struct {
	gorm.Model
	No      string `sql:"-" json:"id"`
	Logo    string `sql:"type:varchar(255)" description:"logo图标" json:"logo"`
	Name    string `sql:"type:varchar(100)" description:"支付模版" json:"name" validate:"required"`
	Method  string `sql:"type:varchar(100)" description:"调用方法名" json:"method" validate:"required"`
	Status  Status `sql:"type:integer;default(1)" description:"状态" json:"status" validate:"required"`
	SiteUrl string `sql:"type:varchar(255)" description:"官网网址" json:"site_url"`
}

func (p *Payment) transform() {
	p.No = strconv.Itoa(int(p.ID))
}

func (p *Payment) AfterSave() error {
	p.transform()
	return nil
}

func (p *Payment) AfterFind() error {
	p.transform()
	return nil
}

type PaymentInstall struct {
	gorm.Model
	No        string  `sql:"-" json:"id"`
	TenantId  string  `sql:"type:char(20)" description:"租户ID" json:"tenant_id" validate:"required"`
	PayId     string  `sql:"type:char(20);index" json:"pay_id" validate:"required"`
	Pay       Payment `gorm:"ForeignKey:PayId;save_associations:false" json:"pay"`
	AppKey    string  `sql:"type:varchar(50)" description:"app key 三方" json:"app_key"`
	AppSecret string  `sql:"type:varchar(100)" description:"app secret 三方" json:"app_secret"`
	Status    Status  `sql:"type:integer;default(1)" description:"状态" json:"status" validate:"required"`
}
