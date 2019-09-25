package pkg

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogger() {
	//todo 后期对接日志系统
	log.SetOutput(os.Stdout)

	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
}
