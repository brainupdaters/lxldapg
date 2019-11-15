package lib

import (
	lfshook "github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LoggingConfig struct {
	Level string
	File  string
}

func SetLoggingConfigDefaults(app string) {
	viper.SetDefault("logging.level", "trace")
	viper.SetDefault("logging.file", app+".log")
}

func InitLogger(cfg LoggingConfig) {
	level, err := log.ParseLevel(cfg.Level)
	if err != nil {
		level = log.TraceLevel
	}
	log.SetLevel(level)

	pathMap := lfshook.PathMap{
		log.TraceLevel: cfg.File,
		log.DebugLevel: cfg.File,
		log.InfoLevel:  cfg.File,
		log.WarnLevel:  cfg.File,
		log.ErrorLevel: cfg.File,
		log.FatalLevel: cfg.File,
		log.PanicLevel: cfg.File,
	}

	log.AddHook(lfshook.NewHook(
		pathMap,
		&log.TextFormatter{},
	))
}
