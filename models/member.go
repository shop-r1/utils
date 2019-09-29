package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shop-r1/utils/pkg"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Member struct {
	gorm.Model
	RecordActive     `sql:"-"`
	ReferrerId       int
	No               string      `sql:"-" json:"id"`
	Referrer         *Member     `gorm:"ForeignKey:ReferrerId" json:"referrer,omitempty"`
	TenantId         string      `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	Username         string      `sql:"type:varchar(100);index" description:"用户名" json:"username"`
	Nickname         string      `sql:"type:varchar(100);index" description:"昵称" json:"nickname"`
	Phone            string      `sql:"type:char(20)" description:"手机号" json:"phone"`
	Description      string      `sql:"type:text" description:"描述" json:"description"`
	OpenId           string      `sql:"type:varchar(100)" description:"三方登录openid" json:"open_id"`
	PasswordHash     string      `sql:"type:varchar(100)" description:"密码hash值" json:"-"`
	Salt             string      `sql:"type:varchar(20)" description:"加盐值" json:"-"`
	RestPasswordHash string      `sql:"type:varchar(100)" description:"重置密码hash值" json:"-"`
	Status           Status      `sql:"default(1)" description:"状态: 1 启用, 0 禁用" json:"status"`
	Password         string      `sql:"-"` //用于暂存密码
	HeadImage        string      `sql:"type:text" description:"头像" json:"head_image"`
	Metadata         []byte      `sql:"json" description:"附加信息" json:"-"`
	Meta             interface{} `sql:"-" json:"meta"`
}

//type SearchMember struct {
//	List      []Member `json:"list"`
//	Total     int      `json:"total"`
//	Page      int      `json:"page"`
//	TotalPage int      `json:"total_page"`
//	Limit     int      `json:"limit"`
//}

func (u *Member) BeforeCreate() (err error) {
	u.generateSalt()
	u.SetPassword(u.Password)
	return nil
}

func (u *Member) BeforeUpdate() (err error) {
	if u.Password != "" {
		//修改密码
		u.generateSalt()
		u.SetPassword(u.Password)
	}
	return err
}

func (u *Member) AfterFind() (err error) {
	u.No = strconv.Itoa(int(u.ID))
	return nil
}

func (u *Member) AfterSave() error {
	u.No = strconv.Itoa(int(u.ID))
	return nil
}

//设置密码
func (u *Member) SetPassword(value string) {
	u.Password = value
	u.generateSalt()
	u.PasswordHash, _ = pkg.SetPassword(u.Password, u.Salt)
}

//获取密码hash
func (u *Member) GetPasswordHash() string {
	passwordHash, err := pkg.SetPassword(u.Password, u.Salt)
	if err != nil {
		return ""
	}
	return passwordHash
}

//生成加盐值
func (u *Member) generateSalt() {
	u.Salt = pkg.GenerateRandomKey16()
}

//验证密码
func (u *Member) Verify() bool {
	Db.Where("username= ?", u.Username).
		Where("tenant_id = ?", u.TenantId).
		Where("status = ?", Enable).First(u)
	return u.GetPasswordHash() == u.PasswordHash
}

//验证是否存在重复username
func (u *Member) ExistUsername() (exist bool, err error) {
	var count int
	err = Db.Model(&Member{}).Where("username = ?", u.Username).
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
