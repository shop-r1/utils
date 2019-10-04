package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

//功能圈
type FunctionCircle struct {
	gorm.Model
	No       string  `sql:"-" json:"id"`
	TenantId string  `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	Name     string  `sql:"type:varchar(100)" description:"名称" json:"name"`
	Album    []Media `gorm:"ForeignKey:FunctionCircleId" description:"相册" json:"album"`
	Content  string  `sql:"type:text" description:"内容" json:"content"`
}

type SearchFunctionCircle struct {
	List      []FunctionCircle `json:"list"`
	Total     int              `json:"total"`
	Page      int              `json:"page"`
	TotalPage int              `json:"total_page"`
	Limit     int              `json:"limit"`
}

//媒体资料
type Media struct {
	CreatedAt        time.Time
	ID               uint    `gorm:"primary_key"`
	No               string  `sql:"-" json:"id"`
	TenantId         string  `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	FunctionCircleId string  `sql:"type:char(20);index" description:"功能圈ID" json:"function_circle_id"`
	BgColor          string  `sql:"type:char(50)" description:"背景颜色" json:"bg_color"`
	BgImage          string  `sql:"type:varchar(255)" description:"背景图片" json:"bg_image"`
	Media            string  `sql:"type:varchar(255)" description:"媒体地址" json:"media"`
	Width            float32 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"宽" json:"width"`
	Height           float32 `sql:"type:DECIMAL(10, 2);default(0.00)" description:"高" json:"height"`
	Video            bool    `description:"视频" json:"video"`
	LinkType         string  `sql:"type:char(50)" description:"关联类型" json:"link_type"`
	LinkId           string  `sql:"type:char(20)" description:"关联ID" json:"link_id"`
	Url              string  `sql:"type:varchar(255)" description:"链接" json:"url"`
}

func (f *FunctionCircle) AfterSave() error {
	f.transform()
	return nil
}

func (f *FunctionCircle) AfterFind() error {
	f.transform()
	return nil
}

func (f *FunctionCircle) transform() {
	f.No = strconv.Itoa(int(f.ID))
}

func (m *Media) AfterSave() error {
	m.transform()
	return nil
}

func (m *Media) AfterFind() error {
	m.transform()
	return nil
}

func (m *Media) transform() {
	m.No = strconv.Itoa(int(m.ID))
}

func (f *FunctionCircle) BeforeSave(tx *gorm.DB) (err error) {
	err = tx.Where("function_circle_id = ?", f.ID).Delete(&Media{}).Error
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}
	for index, m := range f.Album {
		m.TenantId = f.TenantId
		f.Album[index] = m
	}
	return nil
}
