package models

import (
	"strconv"
	"time"
)

type ShoppingCart struct {
	ID                   uint                   `gorm:"primary_key"`
	No                   string                 `sql:"-" json:"id"`
	TenantId             string                 `sql:"type:char(20)" description:"租户ID" json:"tenant_id"`
	MemberId             string                 `sql:"type:char(20)" description:"客户ID" json:"member_id"`
	GoodsId              string                 `sql:"type:char(20)" description:"商品ID" json:"goods_id" validate:"required"`
	Goods                Goods                  `gorm:"save_associations:false" json:"goods" validate:"-"`
	GoodsSpecificationId string                 `sql:"type:char(20)" description:"商品规格ID" json:"goods_specification_id"`
	GoodsSpecification   GoodsSpecification     `gorm:"save_associations:false" json:"goods_specification" validate:"-"`
	WarehouseId          string                 `sql:"type:char(20)" description:"发货仓ID" json:"warehouse_id" validate:"required"`
	Warehouse            GoodsShippingWarehouse `gorm:"save_associations:false" json:"warehouse" validate:"-"`
	PackSpecification    int                    `sql:"type:integer;default(1)" description:"包装规格(默认1)" json:"pack_specification" validate:"required"`
	Quantity             int                    `sql:"type:integer;default(1)" description:"数量" json:"quantity" validate:"required"`
	Selected             bool                   `description:"选中" json:"selected"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (e *ShoppingCart) AfterFind() error {
	e.transform()
	return nil
}

func (e *ShoppingCart) transform() {
	e.No = strconv.Itoa(int(e.ID))
}
