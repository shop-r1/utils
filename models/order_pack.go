package models

import (
	"time"
)

type SendStatus string

const (
	SendWill        SendStatus = "will"        //待发货
	SendAlready     SendStatus = "already"     //已发货
	SendSynchronize SendStatus = "synchronize" //同步物流
)

//订单包裹，不支持更新
type OrderPack struct {
	CreatedAt  time.Time
	ID         uint       `gorm:"primary_key"`
	No         string     `sql:"-" json:"id"`
	DeletedAt  *time.Time `sql:"index"`
	TenantId   string     `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	OrderId    string     `sql:"type:char(20);index" description:"订单ID" json:"order_id"`
	CourierNo  string     `sql:"type:char(50);index" description:"物流单号" json:"courier_no"`
	CourierFee float32    `sql:"type:DECIMAL(10, 2);default(0.00)" description:"快递费" json:"courier_fee"`
	SendStatus SendStatus `sql:"type:char(20);index" description:"发货状态" json:"send_status"`
}

//订单包裹关联商品，不支持更新
type OrderPackGoods struct {
	CreatedAt    time.Time
	ID           uint       `gorm:"primary_key"`
	No           string     `sql:"-" json:"id"`
	DeletedAt    *time.Time `sql:"index"`
	OrderGoodsId string     `sql:"type:char(20);index" description:"商品关联ID" json:"order_goods_id"`
	Quantity     int        `sql:"type:integer" description:"数量" json:"quantity"`
}
