package models

import "github.com/jinzhu/gorm"

type OrderGoods struct {
	gorm.Model
	No                     string              `sql:"-" json:"id"`
	TenantId               string              `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	OrderId                string              `sql:"type:char(20);index" description:"订单ID" json:"order_id"`
	GoodsId                string              `sql:"type:char(20);index" description:"商品ID" json:"goods_id"`
	GoodsSpecId            string              `sql:"type:char(20);index" description:"商品规格ID" json:"goods_specification_id"`
	Quantity               int                 `sql:"type:integer" description:"数量" json:"quantity"`
	Price                  float32             `sql:"type:DECIMAL(10, 2);default(0.00)" description:"价格" json:"price"`
	GoodsData              []byte              `sql:"type:json" description:"商品快照" json:"-"`
	Goods                  *Goods              `sql:"-" description:"商品快照结构体" json:"goods"`
	GoodsSpecificationData []byte              `sql:"type:json" description:"商品规格快照" json:"-"`
	GoodsSpecification     *GoodsSpecification `sql:"-" description:"商品规格快照结构体" json:"goods_specification"`
}
