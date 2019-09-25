package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Role struct {
	gorm.Model
	TenantId    uint `json:"tenant_id"`
	Tenant      Tenant
	No          string `sql:"-" json:"id"`
	Name        string `gorm:"type:varchar(100)" description:"用户名"`
	Description string `gorm:"type:text" description:"描述"`
	Privilege   []byte `gorm:"type:json" description:"权限范围"`
	Status      Status `gorm:"default(1)" description:"状态: 1 启用, 0 禁用"`
}

type SearchRole struct {
	List      []*Role
	Total     int
	Page      int
	TotalPage int
	Limit     int
}

func (r *Role) BeforeSave() error {
	if len(r.Privilege) == 0 {
		r.Privilege = []byte(`[]`)
	}
	return nil
}

func (r *Role) AfterFind() (err error) {
	r.No = strconv.Itoa(int(r.ID))
	return nil
}

func (r *Role) AfterSave() (err error) {
	r.No = strconv.Itoa(int(r.ID))
	return nil
}

//置为不可用状态
func (r *Role) Disable() error {
	r.Status = Disable
	return Db.Model(r).Update("status").Error
}

//置为可用状态
func (r *Role) Enable() error {
	r.Status = Enable
	return Db.Model(r).Update("status").Error
}

//验证是否存在重复name
func (r *Role) ExistName() (exist bool, err error) {
	var count int
	err = Db.Model(&Role{}).Where("name = ?", r.Name).Where("tenant_id = ?", r.TenantId).Count(&count).Error
	if err != nil {
		err = errors.New("数据库查询出错")
		return
	}
	if exist = count > 0; exist {
		log.Error(fmt.Sprintf("name:%s 已经存在", r.Name))
	}
	return
}
