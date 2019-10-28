package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

//寄件人
type Sender struct {
	gorm.Model
	No       string `sql:"-" json:"id"`
	TenantId string `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	MemberId string `sql:"type:char(20);index" description:"客户ID" json:"member_id"`
	Name     string `sql:"type:varchar(100)" description:"姓名" json:"name" validate:"required"`
	Phone    string `sql:"type:varchar(50)" description:"手机号" json:"phone" validate:"required"`
	Country  string `sql:"type:varchar(50)" description:"国家" json:"country"`
	Province string `sql:"type:varchar(50)" description:"省份" json:"province"`
	City     string `sql:"type:varchar(50)" description:"市" json:"city"`
	Address  string `sql:"type:varchar(255)" description:"详细地址" json:"address"`
	Tag      string `sql:"type:varchar(100)" description:"地址标签" json:"tag"`
	Default  bool   `description:"默认地址" json:"default"`
}

type SearchSender struct {
	List      []Sender `json:"list"`
	Total     int      `json:"total"`
	Page      int      `json:"page"`
	TotalPage int      `json:"total_page"`
	Limit     int      `json:"limit"`
}

func (s *Sender) AfterSave() error {
	s.transform()
	return nil
}

func (s *Sender) AfterFind() error {
	s.transform()
	return nil
}

func (s *Sender) transform() {
	s.No = strconv.Itoa(int(s.ID))
}
