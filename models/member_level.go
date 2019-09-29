package models

import "github.com/jinzhu/gorm"

type MemberLevel struct {
	gorm.Model
	RecordActive `sql:"-"`
	TenantId     string  `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	Name         string  `sql:"type:varchar(100);index" description:"用户名" json:"name"`
	HasMarket    bool    `description:"可以开店" json:"has_market"`
	PayIds       string  `sql:"type:text" description:"可用的支付方式" json:"pay_ids"`
	Ratio        float32 `sql:"type:DECIMAL(10, 2)" description:"价格浮动比例" json:"ratio"`
}

//type SearchMemberLevel struct {
//	List      []MemberLevel `json:"list"`
//	Total     int           `json:"total"`
//	Page      int           `json:"page"`
//	TotalPage int           `json:"total_page"`
//	Limit     int           `json:"limit"`
//}
