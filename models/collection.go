package models

import "github.com/jinzhu/gorm"

//收藏夹
type Collection struct {
	gorm.Model
	No                   string `sql:"-" json:"id"`
	TenantId             string `sql:"type:varchar(20)" description:"租户ID" json:"-"`
	MemberId             string `sql:"type:varchar(20)" description:"客户ID" json:"member_id"`
	GoodsId              string `sql:"type:varchar(20)" description:"商品ID" json:"goods_id"`
	GoodsSpecificationId string `sql:"type:varchar(20)" description:"商品规格ID" json:"goods_specification_id"`
}

type SearchCollection struct {
	List      []Collection `json:"list"`
	Total     int          `json:"total"`
	Page      int          `json:"page"`
	TotalPage int          `json:"total_page"`
	Limit     int          `json:"limit"`
}
