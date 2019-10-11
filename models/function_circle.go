package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//功能圈
type FunctionCircle struct {
	gorm.Model
	No       string  `sql:"-" json:"id"`
	TenantId string  `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	Type     string  `sql:"type:char(100);index" description:"类型" json:"type"`
	Name     string  `sql:"type:varchar(100)" description:"名称" json:"name"`
	Media    []Media `gorm:"ForeignKey:FunctionCircleId;save_associations:false" description:"相册" json:"media"`
	Status   Status  `sql:"default(1)" description:"状态" json:"status" validate:"required"`
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
	gorm.Model
	No               string `sql:"-" json:"id"`
	TenantId         string `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	FunctionCircleId string `sql:"type:char(20);index" description:"功能圈ID" json:"function_circle_id"`
	BgColor          string `sql:"type:char(50)" description:"背景颜色" json:"bg_color"`
	BgImage          string `sql:"type:varchar(255)" description:"背景图片" json:"bg_image"`
	Media            string `sql:"type:varchar(255)" description:"媒体地址" json:"media"`
	Video            bool   `description:"视频" json:"video"`
	LinkType         string `sql:"type:char(50)" description:"关联类型" json:"link_type"`
	LinkId           string `sql:"type:char(20)" description:"关联ID" json:"link_id"`
	Url              string `sql:"type:varchar(255)" description:"链接" json:"url"`
	Sort             int    `description:"排序" json:"sort"`
}

//func (f *FunctionCircle) AfterSave() error {
//	f.transform()
//	return nil
//}

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

func (f *FunctionCircle) AfterSave(tx *gorm.DB) (err error) {
	err = tx.Where("function_circle_id = ?", f.ID).Delete(&Media{}).Error
	if err != nil {
		log.Error(err)
		return err
	}
	for index, m := range f.Media {
		m.TenantId = f.TenantId
		m.FunctionCircleId = strconv.Itoa(int(f.ID))
		err = tx.Create(&m).Error
		if err != nil {
			log.Error(err)
			return err
		}
		f.Media[index] = m
	}
	f.transform()
	return nil
}
