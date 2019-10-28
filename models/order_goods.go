package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
)

type OrderGoods struct {
	gorm.Model
	No                     string              `sql:"-" json:"id"`
	OrderId                string              `sql:"type:char(20);index" description:"订单ID" json:"order_id"`
	GoodsId                string              `sql:"type:char(20);index" description:"商品ID" json:"goods_id"`
	GoodsSpecId            string              `sql:"type:char(20);index" description:"商品规格ID" json:"goods_specification_id"`
	Quantity               int                 `sql:"type:integer" description:"数量" json:"quantity"`
	Price                  float32             `sql:"type:DECIMAL(10, 2);default(0.00)" description:"价格" json:"price"`
	GoodsData              []byte              `sql:"type:json" description:"商品快照" json:"-"`
	Goods                  *Goods              `sql:"-" description:"商品快照结构体" json:"goods"`
	GoodsSpecificationData []byte              `sql:"type:json" description:"商品规格快照" json:"-"`
	GoodsSpecification     *GoodsSpecification `sql:"-" description:"商品规格快照结构体" json:"goods_specification"`
	PackSpecification      int                 `description:"包装规格" json:"pack_specification"`
}

func (e *OrderGoods) AfterFind() error {
	e.transform()
	return nil
}

func (e *OrderGoods) BeforeSave() error {
	e.unTransform()
	return nil
}

func (e *OrderGoods) unTransform() {
	e.GoodsData, _ = json.Marshal(e.Goods)
	e.GoodsSpecificationData, _ = json.Marshal(e.GoodsSpecification)
	if e.ID == 0 && e.No != "" && e.No != "0" {
		id, _ := strconv.Atoi(e.No)
		e.ID = uint(id)
	}
}

func (e *OrderGoods) transform() {
	_ = json.Unmarshal(e.GoodsData, &e.Goods)
	_ = json.Unmarshal(e.GoodsSpecificationData, &e.GoodsSpecification)
	e.No = strconv.Itoa(int(e.ID))
}
