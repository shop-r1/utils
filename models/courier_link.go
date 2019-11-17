package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type CourierLink struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	No          string          `sql:"-" json:"id"`
	LinkId      int             `gorm:"primary_key" json:"link_id"`
	LeftRuleId  int             `gorm:"primary_key" json:"left_rule_id"`
	RightRuleId int             `gorm:"primary_key" json:"right_rule_id"`
	LeftRule    CourierPackRule `gorm:"save_associations:false"`
	RightRule   CourierPackRule `gorm:"save_associations:false"`
	ObjectIds   string          `sql:"type:text" description:""`
}

type ObjectLinkCourier struct {
	Id      int     `json:"id"`
	RuleIds [][]int `json:"rule_ids"`
}

type GeneratePack struct {
	GetSelf          bool               `json:"get_self"`
	Alias            string             `json:"alias"`
	Image            string             `json:"image"`
	GoodsId          string             `json:"goods_id"`
	ParentCategoryId string             `json:"parent_category_id"`
	CategoryId       string             `json:"category_id"`
	CourierId        string             `json:"courier_id"`
	Weight           int                `json:"weight"`
	Price            float64            `json:"price"`
	Meta             map[string]int     `json:"meta"`
	MetaPrice        map[string]float64 `json:"meta_price"`
	Quantity         int                `json:"quantity"`
}

type OrderUnitPack struct {
	gorm.Model
	No               string         `sql:"-" json:"id"`
	TenantId         string         `sql:"type:char(20);index" description:"租户ID" json:"tenant_id,omitempty"`
	MemberId         string         `sql:"type:char(20)" description:"客户ID" json:"member_id,omitempty"`
	OrderId          string         `sql:"type:char(20);index" description:"订单ID" json:"order_id"`
	Pack             []GeneratePack `sql:"-" json:"pack"`
	PackData         []byte         `sql:"type:json" json:"-"`
	Weight           int            `json:"weight"`
	GoodsPrice       float64        `sql:"type:DECIMAL(10, 2);default(0.00)" description:"商品价格" json:"goods_price"`
	CourierPrice     float64        `sql:"type:DECIMAL(10, 2);default(0.00)" description:"运费" json:"courier_price"`
	Currency         Currency       `json:"currency"`
	CourierId        string         `sql:"type:char(20);index" description:"物流ID" json:"courier_id"`
	CourierInstallId string         `sql:"type:char(20)" description:"安装的物流ID" json:"courier_install_id"`
	CourierName      string         `sql:"type:varchar(100)" description:"物流名称" json:"courier_name"`
	CourierNo        string         `sql:"type:varchar(100)" description:"物流单号" json:"courier_no"`
	Method           string         `sql:"type:char(20)" description:"物流方法" json:"method"`
	Remark           string         `sql:"type:text" description:"备注" json:"remark"`
	SendStatus       SendStatus     `sql:"type:char(20);index" description:"发货状态" json:"send_status"`
}

func (e *OrderUnitPack) BeforeSave() error {
	e.unTransform()
	return nil
}

func (e *OrderUnitPack) AfterFind() error {
	e.transform()
	return nil
}

func (e *OrderUnitPack) transform() {
	e.No = strconv.Itoa(int(e.ID))
	_ = json.Unmarshal(e.PackData, &e.Pack)
}

func (e *OrderUnitPack) unTransform() {
	e.PackData, _ = json.Marshal(e.Pack)
}
