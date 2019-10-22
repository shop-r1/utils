package models

import (
	"time"
)

//收藏夹
type Collection struct {
	TenantId  string `sql:"type:varchar(20);primary_key" description:"租户ID" json:"-"`
	MemberId  string `sql:"type:varchar(20);primary_key" description:"客户ID" json:"member_id"`
	GoodsId   string `sql:"type:varchar(20);primary_key" description:"商品ID" json:"goods_id"`
	Goods     Goods  `gorm:"save_associations:false" json:"goods" validate:"-"`
	CreatedAt time.Time
}

type SearchCollection struct {
	List      []Collection `json:"list"`
	Total     int          `json:"total"`
	Page      int          `json:"page"`
	TotalPage int          `json:"total_page"`
	Limit     int          `json:"limit"`
}
