package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	TenantId   string `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	MemberId   string `sql:"type:char(20)" description:"客户ID" json:"member_id"`
	CourierId  string `sql:"type:char(20)" description:"物流ID" json:"courier_id"`
	OrderPayId string `sql:"type:char(20)" description:"支付订单ID" json:"order_pay_id"`
	Overage    int    `sql:"type:integer;default(0)" description:"使用余额" json:"overage"`
	Gold       int    `sql:"type:integer;default(0)" description:"金币" json:"gold"`
	Money      int    `sql:"type:integer;default(0)" description:"付款金额" json:"money"`
}

//订单操作记录
type OrderOperateLog struct {
	ID         uint        `gorm:"primary_key"`
	OrderId    string      `sql:"type:char(20);index" description:"订单ID" json:"order_id"`
	Operation  string      `sql:"type:varchar(50)" description:"操作行为" json:"operation"`
	Remark     string      `sql:"type:text" description:"备注" json:"remark"`
	ObjectData []byte      `sql:"type:json" description:"数据状态"`
	Object     interface{} `sql:"-" json:"object"`
	CreatedAt  time.Time
}
