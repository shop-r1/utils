package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type Brand struct {
	gorm.Model
	No          string `sql:"-" json:"id"`
	NameZh      string `sql:"type:varchar(100)" description:"中文名称" json:"name_zh" validate:"required"`
	NameEn      string `sql:"type:varchar(100)" description:"英文名称" json:"name_en" validate:"required"`
	Logo        string `sql:"type:text" description:"logo" json:"logo"`
	SiteUrl     string `sql:"type:varchar(255)" description:"品牌网址" json:"site_url"`
	IndexImg    string `sql:"type:varchar(255)" description:"大图专区" json:"index_img"`
	BgImg       string `sql:"type:varchar(255)" description:"背景图" json:"bg_img"`
	Description string `sql:"type:text" description:"描述" json:"description"`
	Sort        int    `sql:"integer;default(0)" description:"排序" json:"sort"`
	Status      Status `sql:"integer;default(1)" description:"状态" json:"status" validate:"required"`
}

type SearchBrand struct {
	List      []Brand `json:"list"`
	Total     int     `json:"total"`
	Page      int     `json:"page"`
	TotalPage int     `json:"total_page"`
	Limit     int     `json:"limit"`
}

func (b *Brand) AfterFind() error {
	b.No = strconv.Itoa(int(b.ID))
	return nil
}

func (b *Brand) AfterSave() error {
	b.No = strconv.Itoa(int(b.ID))
	return nil
}

func (b *Brand) BeforeSave() error {
	if b.No != "" {
		var id int
		id, _ = strconv.Atoi(b.No)
		b.ID = uint(id)
	}
	return nil
}
