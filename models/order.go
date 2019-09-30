package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	TenantId   string `sql:"type:char(20)" description:"租户ID" json:"tenant_id"`
	MemberId   string `sql:"type:char(20)" description:"客户ID" json:"member_id"`
	CourierId  string `sql:"type:char(20)" description:"物流ID" json:"courier_id"`
	PayOrderId string `sql:"type:char(20)" description:"支付订单ID" json:"pay_order_id"`
}
