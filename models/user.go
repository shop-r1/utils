package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shop-r1/utils/pkg"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type User struct {
	gorm.Model
	TenantId         string `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	Tenant           Tenant `gorm:"save_associations:false" json:"tenant" validate:"-"`
	RoleId           string `sql:"type:char(20);index" description:"角色ID" json:"role_id"`
	Role             Role   `gorm:"save_associations:false" json:"role" validate:"-"`
	No               string `sql:"-" json:"id"`
	Username         string `gorm:"type:varchar(100)" description:"用户名"`
	Nickname         string `gorm:"type:varchar(100)" description:"昵称"`
	Description      string `gorm:"type:text" description:"描述"`
	OpenId           string `gorm:"type:varchar(100)" description:"三方登录openid"`
	PasswordHash     string `gorm:"type:varchar(100)" description:"密码hash值" json:"-"`
	Salt             string `gorm:"type:varchar(20)" description:"加盐值" json:"-"`
	RestPasswordHash string `gorm:"type:varchar(100)" description:"重置密码hash值"`
	Status           Status `gorm:"default(1)" description:"状态: 1 启用, 0 禁用"`
	Password         string `gorm:"-"` //用于暂存密码
	HeadImage        string `gorm:"type:text" description:"头像"`
}

type SearchUser struct {
	List      []*User
	Total     int
	Page      int
	TotalPage int
	Limit     int
}

func (u *User) BeforeCreate() (err error) {
	u.generateSalt()
	u.SetPassword(u.Password)
	return nil
}

func (u *User) BeforeUpdate() (err error) {
	if u.Password != "" {
		//修改密码
		u.generateSalt()
		u.SetPassword(u.Password)
	}
	return err
}

func (u *User) AfterFind() (err error) {
	u.No = strconv.Itoa(int(u.ID))
	return nil
}

func (u *User) AfterSave() error {
	u.No = strconv.Itoa(int(u.ID))
	return nil
}

//置为不可用状态
//func (u *User) Disable() error {
//	u.Status = Disable
//	return Db.Model(u).Update("status").Error
//}

//置为可用状态
//func (u *User) Enable() error {
//	u.Status = Enable
//	return Db.Model(u).Update("status").Error
//}

//设置密码
func (u *User) SetPassword(value string) {
	u.Password = value
	u.generateSalt()
	u.PasswordHash, _ = pkg.SetPassword(u.Password, u.Salt)
}

//获取密码hash
func (u *User) GetPasswordHash() string {
	passwordHash, err := pkg.SetPassword(u.Password, u.Salt)
	if err != nil {
		return ""
	}
	return passwordHash
}

//生成加盐值
func (u *User) generateSalt() {
	u.Salt = pkg.GenerateRandomKey16()
}

//验证密码
func (u *User) Verify() bool {
	Db.Preload("Role.Tenant").
		Where("username= ?", u.Username).
		Where("tenant_id = ?", u.TenantId).
		Where("status = ?", Enable).First(u)
	return u.GetPasswordHash() == u.PasswordHash
}

//验证是否存在重复username
func (u *User) ExistUsername() (exist bool, err error) {
	var count int
	err = Db.Model(&User{}).Where("username = ?", u.Username).
		Where("tenant_id = ?", u.TenantId).Count(&count).Error
	if err != nil {
		err = errors.New("数据库查询出错")
		return
	}
	if exist = count > 0; exist {
		log.Error(fmt.Sprintf("username:%s 已经存在", u.Username))
	}
	return
}
