package models

import "github.com/jinzhu/gorm"

type PaymentOrder struct {
	gorm.Model
	TenantId string  `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	MemberId string  `sql:"type:char(20);index" description:"客户ID" json:"member_id"`
	Member   Member  `gorm:"ForeignKey:MemberId;save_associations:false" json:"member"`
	PayId    string  `sql:"type:char(20);index" description:"支付ID" json:"pay_id"`
	Pay      Payment `gorm:"ForeignKey:PayId;save_associations:false" json:"pay"`
	Overage  int     `sql:"type:integer;default(0)" description:"使用余额" json:"overage"`
	Gold     int     `sql:"type:integer;default(0)" description:"金币" json:"gold"`
	Money    int     `sql:"type:integer;default(0)" description:"付款金额" json:"money"`
}
