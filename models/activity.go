package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type ActivityType string

const (
	ActivityFullGift      ActivityType = "FullGift"      //满赠
	ActivityFullReduction ActivityType = "FullReduction" //满减
	ActivityPanicBuying   ActivityType = "PanicBuying"   //秒杀、抢购
	ActivityFreeShipping  ActivityType = "FreeShipping"  //包邮
)

type Activity struct {
	gorm.Model
	Links        []ActivityLink `gorm:"ForeignKey:ActivityId;save_associations:false" json:"links"`
	Name         string         `sql:"type:varchar(100)" description:"活动名称" json:"name"`
	Description  string         `sql:"type:text" description:"描述" json:"description"`
	IndexImg     string         `sql:"type:varchar(255)" description:"大图专区" json:"index_img"`
	BgImg        string         `sql:"type:varchar(255)" description:"背景图" json:"bg_img"`
	Status       Status         `sql:"integer;default(1)" description:"状态" json:"status" validate:"required"`
	Start        time.Time      `sql:"index" description:"开始时间" json:"start"`
	End          time.Time      `sql:"index" description:"结束时间" json:"end"`
	ActivityType ActivityType   `sql:"type:varchar(50);index" description:"活动类型" json:"activity_type"`
	Extend       interface{}    `sql:"-" description:"活动扩展字段" json:"extend"`
	ExtendData   []byte         `sql:"type:json" description:"活动扩展数据字段" json:"extend_data"`
}

type ExtendFullReduction struct {
	EnoughPrice    float64 `description:"满足金额" json:"enough_price"`
	EnoughQuantity int     `description:"满足数量" json:"enough_quantity"`
	Reduce         float64 `description:"减免金额" json:"reduce"`
}

type ExtendFullGift struct {
	EnoughPrice    float64 `description:"满足金额" json:"enough_price"`
	EnoughQuantity int     `description:"满足数量" json:"enough_quantity"`
	Reduce         float64 `description:"减免金额" json:"reduce"`
}

type ActivityLink struct {
	ID           uint         `gorm:"primary_key"`
	ActivityId   int          `sql:"index" description:"活动ID" json:"activity_id"`
	Activity     Activity     `gorm:"save_associations:false" json:"activity" validate:"-"`
	LinkType     LinkType     `sql:"type:varchar(50);index" description:"关联类型" json:"link_type"`
	ActivityType ActivityType `sql:"type:varchar(50);index" description:"活动类型" json:"activity_type"`
	LinkId       int          `sql:"index" description:"关联ID" json:"link_id"`
	CreatedAt    time.Time
}

func (e *Activity) BeforeSave() (err error) {
	e.ExtendData, err = json.Marshal(e.Extend)
	return err
}

func (e *Activity) AfterSave(tx *gorm.DB) (err error) {
	for _, l := range e.Links {
		l.ActivityId = int(e.ID)
		l.ActivityType = e.ActivityType
		err = tx.Create(&l).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Activity) BeforeFind() (err error) {
	err = json.Unmarshal(e.ExtendData, &e.Extend)
	if err != nil {
		return err
	}
	return nil
}
