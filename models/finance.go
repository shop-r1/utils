package models

import (
	"time"
)

type FinanceType int

const (
	Overage FinanceType = 1
	Gold    FinanceType = 2
)

type Finance struct {
	MemberId      string     `gorm:"primary_key;type:char(20);index" description:"会员ID" json:"member_id"`
	TenantId      string     `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	Overage       float64    `sql:"type:DECIMAL(10, 2);default(0.00)" description:"余额" json:"overage"`
	Gold          float64    `sql:"type:DECIMAL(10, 2);default(0.00)" description:"金豆" json:"gold"`
	FreezeOverage float64    `sql:"type:DECIMAL(10, 2);default(0.00)" description:"余额冻结数" json:"freeze_overage"`
	FreezeGold    float64    `sql:"type:DECIMAL(10, 2);default(0.00)" description:"金豆冻结数" json:"freeze_gold"`
	DeletedAt     *time.Time `sql:"index" json:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type FinanceLog struct {
	ID          uint        `gorm:"primary_key"`
	MemberId    string      `gorm:"primary_key;type:char(20);index" description:"会员ID" json:"member_id"`
	TenantId    string      `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	FinanceType FinanceType `sql:"type:integer;index" description:"日志类型 1:余额变动, 2:金豆变动" json:"finance_type"`
	Old         float64     `sql:"type:DECIMAL(10, 2);default(0.00)" description:"变动前数" json:"old"`
	Change      float64     `sql:"type:DECIMAL(10, 2);default(0.00)" description:"变动数" json:"change"`
	Freeze      float64     `sql:"type:DECIMAL(10, 2);default(0.00)" description:"冻结数" json:"freeze"`
	CreatedAt   time.Time
}

type SearchFinanceLog struct {
	List      []FinanceLog `json:"list"`
	Total     int          `json:"total"`
	Page      int          `json:"page"`
	TotalPage int          `json:"total_page"`
	Limit     int          `json:"limit"`
}
