package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

//收件人
type Consignee struct {
	gorm.Model
	No          string  `sql:"-" json:"id"`
	TenantId    string  `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	MemberId    string  `sql:"type:char(20);index" description:"客户ID" json:"member_id"`
	Member      *Member `gorm:"save_associations:false;ForeignKey:MemberId" json:"member,omitempty" validate:"-"`
	Name        string  `sql:"type:varchar(100)" description:"姓名" json:"name" validate:"required"`
	Phone       string  `sql:"type:varchar(50)" description:"手机号" json:"phone" validate:"required"`
	Country     string  `sql:"type:varchar(50)" description:"国家" json:"country" validate:"required"`
	Province    string  `sql:"type:varchar(50)" description:"省份" json:"province" validate:"required"`
	City        string  `sql:"type:varchar(50)" description:"市" json:"city" validate:"required"`
	Address     string  `sql:"type:varchar(255)" description:"详细地址" json:"address" validate:"required"`
	Tag         string  `sql:"type:varchar(100)" description:"地址标签" json:"tag"`
	Default     bool    `description:"默认地址" json:"default"`
	IdCard      string  `sql:"type:varchar(50)" description:"身份证号码" json:"id_card"`
	IdCardFront string  `sql:"type:varchar(255)" description:"身份证正面" json:"id_card_front"`
	IdCardBack  string  `sql:"type:varchar(255)" description:"身份证背面" json:"id_card_back"`
}

type SearchConsignee struct {
	List      []Consignee `json:"list"`
	Total     int         `json:"total"`
	Page      int         `json:"page"`
	TotalPage int         `json:"total_page"`
	Limit     int         `json:"limit"`
}

func (c *Consignee) AfterSave() error {
	c.transform()
	return nil
}

func (c *Consignee) AfterFind() error {
	c.transform()
	return nil
}

func (c *Consignee) transform() {
	c.No = strconv.Itoa(int(c.ID))
}
