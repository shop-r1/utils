package models

import "github.com/jinzhu/gorm"

type Finance struct {
	gorm.Model
	TenantId string  `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	MemberId string  `sql:"type:char(20);index" description:"会员ID" json:"member_id"`
	Overage  float64 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"余额" json:"overage"`
	Gold     float64 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"金豆" json:"gold"`
}
