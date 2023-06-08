package sentinel

import (
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/logging"
)

// Init sentinel
func Init(appName string) {
	logging.ResetGlobalLoggerLevel(logging.WarnLevel)
	conf := config.NewDefaultConfig()
	conf.Sentinel.App.Name = appName
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()

	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatalf("[Sentinel] init err: %v", err)
	}
}
