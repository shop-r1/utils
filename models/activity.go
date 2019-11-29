package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ActivityType string

const (
	ActivityFullGift      ActivityType = "FullGift"      //满赠
	ActivityFullReduction ActivityType = "FullReduction" //满减
	ActivityPanicBuying   ActivityType = "PanicBuying"   //秒杀、抢购
)

type Activity struct {
	gorm.Model
	Name         string       `sql:"type:varchar(100)" description:"活动名称" json:"name"`
	Description  string       `sql:"type:text" description:"描述" json:"description"`
	IndexImg     string       `sql:"type:varchar(255)" description:"大图专区" json:"index_img"`
	BgImg        string       `sql:"type:varchar(255)" description:"背景图" json:"bg_img"`
	Status       Status       `sql:"integer;default(1)" description:"状态" json:"status" validate:"required"`
	Start        time.Time    `sql:"index" description:"开始时间" json:"start"`
	End          time.Time    `sql:"index" description:"结束时间" json:"end"`
	ActivityType ActivityType `sql:"type:varchar(50);index" description:"活动类型" json:"activity_type"`
}

type ActivityLink struct {
	ID         uint     `gorm:"primary_key"`
	ActivityId int      `sql:"index" description:"活动ID" json:"activity_id"`
	Activity   Activity `gorm:"save_associations:false" json:"activity" validate:"-"`
	LinkType   LinkType `sql:"type:varchar(50);index" description:"关联类型" json:"link_type"`
	LinkId     int      `sql:"index" description:"关联ID" json:"link_id"`
	GoodsIds   string   `sql:"type:text" description:"商品ID集合" json:"goods_ids"`
	Enough     float64  `sql:"type:DECIMAL(10, 2)" description:"满足金额" json:"enough"`
	Reduce     float64  `sql:"type:DECIMAL(10, 2)" description:"减免金额" json:"reduce"`
	CreatedAt  time.Time
}
