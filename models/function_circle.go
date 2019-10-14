package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type SearchFunctionCircle struct {
	List      []FunctionCircle `json:"list"`
	Total     int              `json:"total"`
	Page      int              `json:"page"`
	TotalPage int              `json:"total_page"`
	Limit     int              `json:"limit"`
}

//功能圈
type FunctionCircle struct {
	gorm.Model
	No       string `sql:"-" json:"id"`
	Title    string `sql:"type:char(100);index" description:"类型" json:"title"`
	Type     string `sql:"type:char(100);index" description:"类型" json:"type"`
	Status   Status `sql:"default(1)" description:"状态" json:"status"`
	TenantId string `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	BgColor  string `sql:"type:char(50)" description:"背景颜色" json:"bg_color"`
	BgImage  string `sql:"type:varchar(255)" description:"背景图片" json:"bg_image"`
	Media    string `sql:"type:varchar(255)" description:"媒体地址" json:"media"`
	Video    bool   `description:"视频" json:"video"`
	LinkType string `sql:"type:char(50)" description:"关联类型" json:"link_type"`
	LinkId   string `sql:"type:char(20)" description:"关联ID" json:"link_id"`
	Content  string `sql:"type:text" description:"内容" json:"content"`
	Url      string `sql:"type:varchar(255)" description:"链接" json:"url"`
	Sort     int    `description:"排序" json:"sort"`
}

func (f *FunctionCircle) AfterFind() error {
	f.transform()
	return nil
}

func (f *FunctionCircle) transform() {
	f.No = strconv.Itoa(int(f.ID))
}

func (f *FunctionCircle) AfterSave(tx *gorm.DB) (err error) {
	f.transform()
	return nil
}
