package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderStatus string

const (
	OrderCreate    OrderStatus = "created"   //创建订单,提交订单
	OrderPending   OrderStatus = "pending"   //待支付
	OrderAlready   OrderStatus = "already"   //已支付
	OrderShipping  OrderStatus = "shipping"  //已经发货
	OrderRefund    OrderStatus = "refund"    //退款
	OrderCanceled  OrderStatus = "canceled"  //取消
	OrderCompleted OrderStatus = "completed" //已完成
)

type Order struct {
	gorm.Model
	No            string       `sql:"-" json:"id"`
	TenantId      string       `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	MemberId      string       `sql:"type:char(20)" description:"客户ID" json:"member_id"`
	CourierId     string       `sql:"type:char(20)" description:"物流ID" json:"courier_id"`
	OrderPayId    string       `sql:"type:char(20)" description:"支付订单ID" json:"order_pay_id"`
	OrderStatus   OrderStatus  `sql:"type:char(20);index" description:"订单状态" json:"order_status"`
	ConsigneeId   string       `sql:"type:char(20);index" description:"收件人ID" json:"consignee_id"`
	ConsigneeData []byte       `sql:"type:json" description:"收件人快照" json:"-"`
	Consignee     Consignee    `sql:"-" description:"收件人快照结构体" json:"consignee"`
	SenderId      string       `sql:"type:char(20);index" description:"寄件人ID" json:"sender_id"`
	SenderData    []byte       `sql:"type:json" description:"寄件人快照" json:"-"`
	Sender        Sender       `sql:"-" description:"寄件人快照结构体" json:"sender"`
	Money         float32      `sql:"type:DECIMAL(10, 2);default(0.00)" description:"付款总金额" json:"money"`
	Currency      Currency     `sql:"type:varchar(20)" description:"币种" json:"currency"`
	Overage       float32      `sql:"type:DECIMAL(10, 2);default(0.00)" description:"使用余额" json:"overage"`
	Gold          float32      `sql:"type:DECIMAL(10, 2);default(0.00)" description:"金币" json:"gold"`
	CourierFee    float32      `sql:"type:DECIMAL(10, 2);default(0.00)" description:"快递费" json:"courier_fee"`
	ReduceFee     float32      `sql:"type:DECIMAL(10, 2);default(0.00)" description:"减免金额" json:"reduce_fee"`
	WarehouseId   string       `sql:"type:char(20);index" description:"发货仓ID" json:"warehouse_id"`
	Goods         []OrderGoods `gorm:"ForeignKey:OrderId" description:"商品关联" json:"goods"`
	Packs         []OrderPack  `gorm:"ForeignKey:OrderId;save_associations:false" description:"包裹关联" json:"packs"`
}

type SearchOrder struct {
	List      []Order `json:"list"`
	Total     int     `json:"total"`
	Page      int     `json:"page"`
	TotalPage int     `json:"total_page"`
	Limit     int     `json:"limit"`
}

//订单操作记录
type OrderOperateLog struct {
	ID         uint        `gorm:"primary_key"`
	No         string      `sql:"-" json:"id"`
	TenantId   string      `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	OrderId    string      `sql:"type:char(20);index" description:"订单ID" json:"order_id"`
	Operation  string      `sql:"type:varchar(50)" description:"操作行为" json:"operation"`
	Remark     string      `sql:"type:text" description:"备注" json:"remark"`
	ObjectData []byte      `sql:"type:json" description:"数据状态"`
	Object     interface{} `sql:"-" json:"object"`
	CreatedAt  time.Time
}

type SearchOrderOperateLog struct {
	List      []OrderOperateLog `json:"list"`
	Total     int               `json:"total"`
	Page      int               `json:"page"`
	TotalPage int               `json:"total_page"`
	Limit     int               `json:"limit"`
}
