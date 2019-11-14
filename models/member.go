package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shop-r1/utils/pkg"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Member struct {
	gorm.Model
	TenantId         string      `sql:"type:char(20);index" description:"租户ID" json:"tenant_id"`
	Region           string      `sql:"type:varchar(100)" description:"地区" json:"region"`
	ReferrerId       string      `json:"referrer_id"`
	No               string      `sql:"-" json:"id"`
	LevelId          string      `sql:"type:char(20);index" description:"客户等级ID" json:"level_id" validate:"-"`
	Level            MemberLevel `gorm:"save_associations:false;ForeignKey:LevelId" json:"level" validate:"-"`
	Referrer         *Member     `gorm:"save_associations:false;ForeignKey:ReferrerId" json:"referrer,omitempty" validate:"-"`
	Username         string      `sql:"type:varchar(100);index" description:"用户名" json:"username"`
	Nickname         string      `sql:"type:varchar(100);index" description:"昵称" json:"nickname"`
	Phone            string      `sql:"type:char(20)" description:"手机号" json:"phone"`
	Description      string      `sql:"type:text" description:"描述" json:"description"`
	OpenId           string      `sql:"type:varchar(100)" description:"三方登录openid" json:"open_id"`
	UnionId          string      `sql:"type:varchar(100)" description:"三方登录unionid" json:"union_id"`
	PasswordHash     string      `sql:"type:varchar(100)" description:"密码hash值" json:"-"`
	Salt             string      `sql:"type:varchar(20)" description:"加盐值" json:"-"`
	RestPasswordHash string      `sql:"type:varchar(100)" description:"重置密码hash值" json:"-"`
	Status           Status      `sql:"default(1)" description:"状态: 1 启用, 0 禁用" json:"status"`
	Password         string      `sql:"-" json:"password,omitempty"` //用于暂存密码
	HeadImage        string      `sql:"type:text" description:"头像" json:"head_image"`
	Metadata         []byte      `sql:"json" description:"附加信息" json:"-"`
	Meta             interface{} `sql:"-" json:"meta"`
	Address          string      `sql:"type:varchar(100)" description:"用户地址" json:"address"`
	Sex              int         `json:"sex"` // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
}

type SearchMember struct {
	List      []Member `json:"list"`
	Total     int      `json:"total"`
	Page      int      `json:"page"`
	TotalPage int      `json:"total_page"`
	Limit     int      `json:"limit"`
}

func (u *Member) BeforeSave() (err error) {
	if u.Password != "" {
		u.generateSalt()
		u.SetPassword(u.Password)
	}
	return nil
}

func (u *Member) AfterSave() error {
	u.No = strconv.Itoa(int(u.ID))
	u.Password = ""
	return nil
}

func (u *Member) AfterFind() error {
	u.transform()
	return nil
}

func (u *Member) transform() {
	u.No = strconv.Itoa(int(u.ID))
	if len(u.Metadata) > 0 {
		_ = json.Unmarshal(u.Metadata, &u.Meta)
	}
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
