package tools

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	log "github.com/sirupsen/logrus"
)

var (
	Conf map[string]interface{}
)

func InitConf() {
	var err error
	err = config.Load(
		env.NewSource(),
		file.NewSource(file.WithPath("conf/app.yml")),
	)
	if err != nil {
		log.Fatal(err)
	}
	Conf = config.Map()
}
