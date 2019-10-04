package models

import "github.com/jinzhu/gorm"

type ShoppingCart struct {
	gorm.Model
	No                   string               `sql:"-" json:"id"`
	TenantId             string               `sql:"type:char(20)" description:"租户ID" json:"tenant_id"`
	MemberId             string               `sql:"type:char(20)" description:"客户ID" json:"member_id" validate:"required"`
	GoodsId              string               `sql:"type:char(20)" description:"商品ID" json:"goods_id" validate:"required"`
	GoodsSpecificationId string               `sql:"type:char(20)" description:"商品规格ID" json:"goods_specification_id"`
	WarehouseId          string               `sql:"type:char(20)" description:"发货仓ID" json:"warehouse_id" validate:"required"`
	PackSpecification    int                  `sql:"type:integer;default(1)" description:"包装规格(默认1)" json:"pack_specification" validate:"required"`
	Quantity             int                  `sql:"type:integer;default(1)" description:"数量" json:"quantity" validate:"required"`
	SnapshotData         []byte               `sql:"type:json" description:"快照" json:"-"`
	Snapshot             ShoppingCartSnapshot `sql:"-" description:"快照结构体" json:"snapshot"`
}

type ShoppingCartSnapshot struct {
	Goods              Goods              `json:"goods"`
	GoodsSpecification GoodsSpecification `json:"goods_specification"`
	//活动快照
}
