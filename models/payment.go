package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type PayType string

const (
	Balance PayType = "balance"
	Online  PayType = "online"
	Voucher PayType = "voucher"
)

type Payment struct {
	gorm.Model
	No          string  `sql:"-" json:"id"`
	Logo        string  `sql:"type:varchar(255)" description:"logo图标" json:"logo"`
	Name        string  `sql:"type:varchar(100)" description:"支付模版" json:"name" validate:"required"`
	Method      string  `sql:"type:varchar(100)" description:"调用方法名" json:"method" validate:"required"`
	Status      Status  `sql:"type:integer;default(1)" description:"状态" json:"status" validate:"required"`
	SiteUrl     string  `sql:"type:varchar(255)" description:"官网网址" json:"site_url"`
	Type        PayType `sql:"type:char(20);index" description:"类型" json:"type"`
	Description string  `sql:"text" description:"描述" json:"description"`
}

type SearchPayment struct {
	List      []Payment `json:"list"`
	Total     int       `json:"total"`
	Page      int       `json:"page"`
	TotalPage int       `json:"total_page"`
	Limit     int       `json:"limit"`
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
	No          string  `sql:"-" json:"id"`
	Used        bool    `description:"安装" json:"used"`
	TenantId    string  `sql:"type:char(20)" description:"租户ID" json:"tenant_id"`
	PaymentId   int     `sql:"type:integer;index" json:"payment_id"`
	Payment     Payment `gorm:"ForeignKey:PaymentId;save_associations:false" json:"payment" validate:"-"`
	AppKey      string  `sql:"type:varchar(50)" description:"app key 三方" json:"app_key"`
	AppSecret   string  `sql:"type:varchar(100)" description:"app secret 三方" json:"app_secret"`
	Image       string  `sql:"type:text" description:"凭证支付所需图片" json:"image"`
	Description string  `sql:"text" description:"描述" json:"description"`
}

func (p *PaymentInstall) transform() {
	p.No = strconv.Itoa(int(p.ID))
}

func (p *PaymentInstall) AfterSave() error {
	p.transform()
	return nil
}

func (p *PaymentInstall) AfterFind() error {
	p.transform()
	return nil
}
