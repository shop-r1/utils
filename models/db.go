package models

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	"github.com/shop-r1/utils/db"
	log "github.com/sirupsen/logrus"
)

type Status int

const (
	_ Status = iota
	Enable
	Disable
)

type Condition struct {
	Limit int
	Page  int
	Where map[string][]interface{}
}

var (
	Db   *gorm.DB
	Conf map[string]interface{}
)

func InitDb(logMode bool, models ...interface{}) {
	var err error
	err = config.Load(
		env.NewSource(),
		file.NewSource(file.WithPath("conf/app.yml")),
	)
	if err != nil {
		log.Fatal(err)
	}
	Conf = config.Map()

	Db, err = db.DbInit(Conf["databaseDriver"].(string), Conf["databaseUrl"].(string))
	if err != nil {
		log.Fatal(err)
	}
	Db.LogMode(logMode)
	Db.AutoMigrate(models...)
}
