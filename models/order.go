package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
)

type OrderStatus string

const (
	OrderCreate     OrderStatus = "created"     //创建订单,提交订单
	OrderNeedVerify OrderStatus = "need-verify" //待验证,自提订单
	OrderPending    OrderStatus = "pending"     //待支付
	OrderAlready    OrderStatus = "already"     //已支付
	OrderShipping   OrderStatus = "shipping"    //已经发货
	OrderRefund     OrderStatus = "refund"      //退款
	OrderCanceled   OrderStatus = "canceled"    //取消
	OrderCompleted  OrderStatus = "completed"   //已完成
)

type Order struct {
	gorm.Model
	No              string               `sql:"-" json:"id"`
	TenantId        string               `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	MemberId        string               `sql:"type:char(20)" description:"客户ID" json:"member_id"`
	CourierId       string               `sql:"type:char(20)" description:"物流ID" json:"courier_id"`
	OrderPayId      string               `sql:"type:char(20)" description:"支付订单ID" json:"order_pay_id"`
	OrderStatus     OrderStatus          `sql:"type:char(20);index" description:"订单状态" json:"order_status"`
	ConsigneeId     string               `sql:"type:char(20);index" description:"收件人ID" json:"consignee_id"`
	ConsigneeData   string               `sql:"type:text" description:"收件人快照" json:"-"`
	Consignee       Consignee            `sql:"-" description:"收件人快照结构体" json:"consignee" validate:"-"`
	SenderId        string               `sql:"type:char(20);index" description:"寄件人ID" json:"sender_id"`
	SenderData      string               `sql:"type:text" description:"寄件人快照" json:"-"`
	Sender          Sender               `sql:"-" description:"寄件人快照结构体" json:"sender" validate:"-"`
	Money           float32              `sql:"type:DECIMAL(10, 2);default(0.00)" description:"付款总金额" json:"money"`
	Currency        Currency             `sql:"type:varchar(20)" description:"币种" json:"currency"`
	Overage         float32              `sql:"type:DECIMAL(10, 2);default(0.00)" description:"使用余额" json:"overage"`
	Gold            float32              `sql:"type:DECIMAL(10, 2);default(0.00)" description:"金币" json:"gold"`
	CourierPrice    float32              `sql:"type:DECIMAL(10, 2);default(0.00)" description:"快递费" json:"courier_price"`
	GoodsPrice      float32              `sql:"type:DECIMAL(10, 2);default(0.00)" description:"商品总价" json:"goods_price"`
	Price           float32              `sql:"type:DECIMAL(10, 2);default(0.00)" description:"订单总价" json:"price_old"`
	PriceMap        map[Currency]float32 `sql:"-" json:"price"`
	ReduceFee       float32              `sql:"type:DECIMAL(10, 2);default(0.00)" description:"减免金额" json:"reduce_fee"`
	WarehouseId     string               `sql:"type:char(20);index" description:"发货仓ID" json:"warehouse_id"`
	OrderGoods      []OrderGoods         `gorm:"ForeignKey:OrderId;save_associations:false" description:"商品关联" json:"goods" validate:"-"`
	OrderUnitPacks  []OrderUnitPack      `gorm:"ForeignKey:OrderId;save_associations:false" description:"包裹关联" json:"packs" validate:"-"`
	GetSelf         bool                 `description:"自提" json:"get_self"`
	PaymentIds      string               `sql:"type:text" description:"可用的支付方式" json:"-"`
	PaymentIdsArray []string             `sql:"-" json:"payment_ids"`
	Remark          string               `sql:"type:text" description:"备注" json:"remark"`
	GoodsName       string               `sql:"type:text" description:"所有商品名称(搜索用)" json:"goods_name"`
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

func (e *Order) BeforeSave() error {
	e.UnTransform()
	return nil
}

func (e *Order) AfterCreate(tx *gorm.DB) (err error) {
	for i, g := range e.OrderGoods {
		g.OrderId = strconv.Itoa(int(e.ID))
		err = tx.Create(&g).Error
		if err != nil {
			return err
		}
		e.OrderGoods[i] = g
	}
	for i, p := range e.OrderUnitPacks {
		p.OrderId = strconv.Itoa(int(e.ID))
		err = tx.Create(&p).Error
		if err != nil {
			return err
		}
		e.OrderUnitPacks[i] = p
	}
	return nil
}

func (e *Order) AfterFind() error {
	e.transform()
	return nil
}

func (e *Order) transform() {
	e.No = strconv.Itoa(int(e.ID))
	_ = json.Unmarshal([]byte(e.ConsigneeData), &e.Consignee)
	_ = json.Unmarshal([]byte(e.SenderData), &e.Sender)
	if e.PaymentIds != "" {
		e.PaymentIdsArray = strings.Split(e.PaymentIds, ",")
	}
}

func (e *Order) UnTransform() {
	var rb []byte
	if e.Consignee.No != "" || e.Consignee.Name != "" {
		id, _ := strconv.Atoi(e.Consignee.No)
		e.Consignee.ID = uint(id)
		rb, _ = json.Marshal(e.Consignee)
		e.ConsigneeData = string(rb)
	}
	rb = make([]byte, 0)
	if e.Sender.No != "" || e.Sender.Name != "" {
		id, _ := strconv.Atoi(e.Sender.No)
		e.Sender.ID = uint(id)
		rb, _ = json.Marshal(e.Sender)
		e.SenderData = string(rb)
	}
	e.PaymentIds = strings.Join(e.PaymentIdsArray, ",")
}
