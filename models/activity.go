package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
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
	No                 string         `sql:"-" json:"id"`
	TenantId           string         `gorm:"primary_key" sql:"type:char(20);index" description:"租户ID" json:"-" `
	Links              []ActivityLink `gorm:"ForeignKey:ActivityId;save_associations:false" json:"links"`
	Name               string         `sql:"type:varchar(100)" description:"活动名称" json:"name"`
	Description        string         `sql:"type:text" description:"描述" json:"description"`
	IndexImg           string         `sql:"type:varchar(255)" description:"大图专区" json:"index_img"`
	BgImg              string         `sql:"type:varchar(255)" description:"背景图" json:"bg_img"`
	Status             Status         `sql:"integer;default(1)" description:"状态" json:"status" validate:"required"`
	StartAt            time.Time      `sql:"index" description:"开始时间" json:"start_at"`
	EndAt              time.Time      `sql:"index" description:"结束时间" json:"end_at"`
	Sort               int            `description:"排序" json:"sort"`
	ActivityType       ActivityType   `sql:"type:varchar(50);index" description:"活动类型" json:"activity_type"`
	Metadata           []byte         `description:"附加信息" json:"-"`
	Meta               interface{}    `sql:"-" description:"附加信息结构" json:"meta"`
	Extend             ExtendActivity `sql:"-" description:"活动扩展字段" json:"extend"`
	ExtendData         []byte         `sql:"type:json" description:"活动扩展数据字段" json:"-"`
	MemberLevelIds     []string       `sql:"-" description:"可参加的客户等级ID集" json:"member_level_ids"`
	MemberLevelIdsData string         `sql:"type:text" description:"可参加的客户等级ID集" json:"-"`
	WarehouseIds       []string       `sql:"-" description:"发货仓ID集" json:"warehouse_ids"`
	WarehouseIdsData   string         `sql:"type:text" description:"发货仓ID集" json:"-"`
}

type ExtendType int

const (
	ExtendTypePrice    ExtendType = 1
	ExtendTypeQuantity ExtendType = 2
)

type ExtendActivity struct {
	ExtendType     ExtendType  `description:"参数类型" json:"extend_type"`
	EnoughPrice    float64     `description:"满足金额" json:"enough_price"`
	EnoughQuantity int         `description:"满足数量" json:"enough_quantity"`
	Reduce         float64     `description:"减免金额" json:"reduce"`
	GiftGoods      []GiftGoods `description:"赠品" json:"gift_goods"`
}

type GiftGoods struct {
	GoodsId  string `description:"商品ID" json:"goods_id"`
	Name     string `description:"名称" json:"name"`
	Image    string `description:"图片" json:"image"`
	Quantity int    `description:"数量" json:"quantity" `
}

type ActivityLink struct {
	ID           uint         `gorm:"primary_key"`
	No           string       `sql:"-" json:"id"`
	TenantId     string       `gorm:"primary_key" sql:"type:char(20);index" description:"租户ID" json:"-" `
	ActivityId   int          `sql:"index" description:"活动ID" json:"activity_id"`
	Activity     *Activity    `gorm:"save_associations:false" json:"activity" validate:"-"`
	LinkType     LinkType     `sql:"type:varchar(50);index" description:"关联类型" json:"link_type"`
	ActivityType ActivityType `sql:"type:varchar(50);index" description:"活动类型" json:"activity_type"`
	LinkId       string       `sql:"type:char(20);index" description:"关联ID" json:"link_id"`
	Name         string       `sql:"type:varchar(100)" description:"名称" json:"name"`
	Image        string       `sql:"type:varchar(255)" description:"图片" json:"image"`
	CreatedAt    time.Time
}

func (e *Activity) BeforeSave() (err error) {
	e.ExtendData, err = json.Marshal(e.Extend)
	if len(e.MemberLevelIds) > 0 {
		e.MemberLevelIdsData = strings.Join(e.MemberLevelIds, ",")
	}
	if len(e.WarehouseIds) > 0 {
		e.WarehouseIdsData = strings.Join(e.WarehouseIds, ",")
	}
	if e.Meta != nil {
		e.Metadata, _ = json.Marshal(e.Meta)
	} else {
		e.Metadata = []byte(`{}`)
	}
	return err
}

func (e *Activity) AfterSave(tx *gorm.DB) (err error) {
	err = tx.Where("activity_id = ?", e.ID).Delete(&ActivityLink{}).Error
	if err != nil {
		return err
	}
	for i, l := range e.Links {
		l.ActivityId = int(e.ID)
		l.ActivityType = e.ActivityType
		l.TenantId = e.TenantId
		err = tx.Create(&l).Error
		if err != nil {
			return err
		}
		e.Links[i] = l
	}
	return nil
}

func (e *Activity) AfterFind() (err error) {
	e.No = strconv.Itoa(int(e.ID))
	if len(e.Metadata) > 0 {
		_ = json.Unmarshal(e.Metadata, &e.Meta)
	}
	err = json.Unmarshal(e.ExtendData, &e.Extend)
	if err != nil {
		return err
	}
	if len(e.MemberLevelIdsData) > 0 {
		e.MemberLevelIds = strings.Split(e.MemberLevelIdsData, ",")
	}
	if len(e.WarehouseIdsData) > 0 {
		e.WarehouseIds = strings.Split(e.WarehouseIdsData, ",")
	}
	return nil
}

func (e *ActivityLink) AfterFind() (err error) {
	e.No = strconv.Itoa(int(e.ID))
	return nil
}

type SearchActivity struct {
	List      []Activity `json:"list"`
	Total     int        `json:"total"`
	Page      int        `json:"page"`
	TotalPage int        `json:"total_page"`
	Limit     int        `json:"limit"`
}
