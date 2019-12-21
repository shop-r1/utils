package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
)

type PaymentStatus string

const (
	PaymentPending PaymentStatus = "pending"
	PaymentSuccess PaymentStatus = "success"
	PaymentAlready PaymentStatus = "already" //客户已付款
	PaymentFailed  PaymentStatus = "failed"
)

type PaymentOrder struct {
	gorm.Model
	No               string            `sql:"-" json:"id"`
	TenantId         string            `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	MemberId         string            `sql:"type:char(20);index" description:"客户ID" json:"member_id"`
	OrderId          string            `sql:"type:char(20);index" description:"订单ID" json:"order_id" validate:"required"`
	PaymentInstallId string            `sql:"type:char(20);index" description:"支付ID" json:"payment_install_id" validate:"required"`
	PaymentInstall   PaymentInstall    `gorm:"ForeignKey:PaymentInstallId;save_associations:false" json:"payment_install" validate:"-"`
	Method           PaymentMethodType `sql:"type:varchar(20)" description:"支付方法" json:"method"`
	Overage          float64           `sql:"type:DECIMAL(10, 2);default(0.00)" description:"使用余额" json:"overage"`
	Currency         Currency          `sql:"type:char(10)" description:"币种" json:"currency"`
	Gold             float64           `sql:"type:DECIMAL(10, 2);default(0.00)" description:"金币" json:"gold"`
	OrderFee         float64           `sql:"type:DECIMAL(10, 2);default(0.00)" description:"订单金额" json:"order_fee"`
	RealFee          float64           `sql:"type:DECIMAL(10, 2);default(0.00)" description:"付款金额" json:"real_fee"`
	AudToCny         float64           `sql:"type:DECIMAL(10, 2);default(0.00)" description:"汇率" json:"aud_to_cny"`
	PayUrl           string            `sql:"type:text" description:"付款链接" json:"pay_url"`
	Redirect         string            `sql:"type:text" description:"付款后跳转" json:"redirect"`
	Token            string            `sql:"type:varchar(100)" description:"回调验证" json:"-"`
	ExternalOrderId  string            `sql:"type:varchar(100)" description:"三方支付订单ID" json:"external_order_id"`
	PaymentStatus    PaymentStatus     `sql:"type:char(10);index" description:"付款状态" json:"payment_status"`
	NotifyUrl        string            `sql:"-" json:"notify_url,omitempty"`
	VoucherData      []byte            `sql:"type:json" description:"付款凭证" json:"-"`
	VoucherImages    []string          `sql:"-" description:"付款凭证" json:"voucher_images"`
	Remark           string            `sql:"type:text" description:"付款备注" json:"remark"`
	Params           interface{}       `sql:"-" json:"params"`
	AppId            string            `sql:"-" json:"app_id,omitempty"`
}

func (e *PaymentOrder) BeforeCreate() error {
	e.VoucherData = []byte(`[]`)
	return nil
}

func (e *PaymentOrder) BeforeUpdate() error {
	e.unTransform()
	return nil
}

func (e *PaymentOrder) AfterCreate() error {
	e.No = strconv.Itoa(int(e.ID))
	return nil
}

func (e *PaymentOrder) AfterFind() error {
	e.transform()
	return nil
}

func (e *PaymentOrder) transform() {
	e.No = strconv.Itoa(int(e.ID))
	_ = json.Unmarshal(e.VoucherData, &e.VoucherImages)
}

func (e *PaymentOrder) unTransform() {
	if len(e.VoucherImages) > 0 {
		e.VoucherData, _ = json.Marshal(e.VoucherImages)
	} else {
		e.VoucherData = []byte(`[]`)
	}
}
