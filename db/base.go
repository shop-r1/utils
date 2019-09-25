package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

//通用更新方法
func Update(orm *gorm.DB, object interface{}) error {
	if orm == nil {
		log.Error("orm db is empty")
		return errors.New("数据库连接对象为空")
	}
	err := orm.Model(object).Update(object).Error
	if err != nil {
		log.Error(err)
		err = errors.New("数据库操作错误，更新失败")
	}
	return err
}

//通用创建方法
func Create(orm *gorm.DB, object interface{}) error {
	if orm == nil {
		log.Error("orm db is empty")
		return errors.New("数据库连接对象为空")
	}
	err := orm.Create(object).Error
	if err != nil {
		log.Error(err)
		err = errors.New("数据库操作错误，创建失败")
	}
	return err
}

func Delete(orm *gorm.DB, object interface{}, id uint64) error {
	return delete(orm, object, "id = ?", id)
}

func delete(orm *gorm.DB, object interface{}, where ...interface{}) error {
	if orm == nil {
		log.Error("orm db is empty")
		return errors.New("数据库连接对象为空")
	}
	err := orm.Delete(object, where...).Error
	if err != nil {
		log.Error(err)
		err = errors.New("数据库操作错误, 删除失败")
	}
	return err
}

func DeleteBatchByIds(orm *gorm.DB, object interface{}, ids ...uint64) error {
	return delete(orm, object, "id IN (?)", ids)
}

func Read(orm *gorm.DB, object interface{}, id uint) error {
	if orm == nil {
		log.Error("orm db is empty")
		return errors.New("数据库连接对象为空")
	}
	err := orm.First(object, id).Error
	if err != nil {
		log.Error(err)
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("未查找到数据")
		}
		err = errors.New("数据库操作错误, 查询失败")
	}
	return err
}

func FindById(orm *gorm.DB, id uint, object interface{}) (err error) {
	if orm == nil {
		log.Error("orm db is empty")
		object = nil
		return errors.New("数据库连接对象为空")
	}
	err = orm.First(object, id).Error
	if err != nil {
		log.Error(err)
		if gorm.IsRecordNotFoundError(err) {
			object = nil
			err = nil
		} else {
			err = errors.New("数据库查询错误")
		}
	}
	return
}

//通用查询方法
func Search(orm *gorm.DB, object interface{}, list interface{}, condition map[string][]interface{}, count *int, limit, offset int, preloads ...string) error {
	var err error
	qs := orm.Model(object)
	for key, args := range condition {
		if len(args) == 0 {
			qs = qs.Order(key)
		} else {
			qs = qs.Where(key, args...)
		}
	}
	for _, preload := range preloads {
		qs = qs.Preload(preload)
	}
	if err = qs.Limit(limit).Offset(offset).Find(list).Limit(-1).Offset(-1).Count(count).Error; err != nil {
		log.Error(err)
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return errors.New("数据库查询错误")
	}
	return nil
}
