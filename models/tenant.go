package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shop-r1/utils/pkg"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Tenant struct {
	gorm.Model
	No             string    `sql:"-" json:"id"`
	Name           string    `gorm:"type:varchar(100)" description:"名称"`
	Contact        string    `gorm:"type:varchar(20)" description:"联系方式"`
	Description    string    `gorm:"type:text" description:"描述"`
	Secret         string    `gorm:"type:varchar(100)" description:"密钥"`
	System         bool      `gorm:"type:boolean" description:"是否为租户系统用户"`
	Expired        time.Time `gorm:"index:idx_expired" description:"到期时间"`
	Status         Status    `gorm:"default(1);index:idx_status" description:"状态: 1 启用, 2 禁用"`
	Domain         string    `gorm:"type:varchar(255);unique" description:"租户独立域名"`
	GenerateSecret bool      `gorm:"-"`
}

type SearchTenant struct {
	List      []*Tenant
	Total     int
	Page      int
	TotalPage int
	Limit     int
}

//自动开户，平台租户系统
func Init自动开户(t *Tenant, r *Role, u *User) {
	var err error
	t = &Tenant{
		Name:           "租户平台管理",
		Description:    "租户管理平台初始化账号",
		System:         true,
		Expired:        time.Now().Add(8760 * time.Hour), //有效期一年
		Domain:         Conf["tenantDomain"].(string),
		GenerateSecret: true,
		Status:         Enable,
	}
	tx := Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			panic("初始化失败")
		}
		tx.Commit()
	}()
	if err = tx.Create(t).Error; err != nil {
		log.Error(err)
		return
	}
	r = &Role{
		TenantId:    t.ID,
		Name:        "管理员",
		Description: "租户管理平台初始化角色",
		Status:      Enable,
	}
	if err = tx.Create(r).Error; err != nil {
		log.Error(err)
		return
	}
	u = &User{
		TenantId:    t.ID,
		RoleId:      r.ID,
		Username:    "admin",
		Nickname:    "admin",
		Description: "租户管理平台初始化用户",
		Password:    "shop-r12345678",
		Status:      Enable,
	}
	if err = tx.Create(u).Error; err != nil {
		log.Error(err)
	}
	return
}

func (t *Tenant) BeforeCreate() (err error) {
	if t.Expired.Unix() < 0 {
		t.Expired = time.Now()
	}
	t.Secret = pkg.GenerateRandomKey20()
	return err
}

func (t *Tenant) BeforeUpdate() (err error) {
	if t.GenerateSecret {
		t.Secret = pkg.GenerateRandomKey20()
	}
	return err
}

func (t *Tenant) AfterFind() (err error) {
	t.No = strconv.Itoa(int(t.ID))
	return nil
}

func (t *Tenant) AfterSave() (err error) {
	t.No = strconv.Itoa(int(t.ID))
	return nil
}

//验证是否存在重复client_id
func (t *Tenant) ExistName() (exist bool, err error) {
	var count int
	err = Db.Model(&Tenant{}).Where("name = ?", t.Name).Count(&count).Error
	if err != nil {
		err = errors.New("数据库查询出错")
		return
	}
	if exist = count > 0; exist {
		log.Error(fmt.Sprintf("name:%s 已经存在", t.Name))
	}
	return
}

//验证是否存在重复domain
func (t *Tenant) ExistDomain() (exist bool, err error) {
	var count int
	err = Db.Model(&Tenant{}).Where("domain = ?", t.Domain).Count(&count).Error
	if err != nil {
		log.Error(err)
		err = errors.New("数据库查询出错")
		return
	}
	if exist = count > 0; exist {
		log.Error(fmt.Sprintf("domain:%s 已经存在", t.Domain))
	}
	return
}
