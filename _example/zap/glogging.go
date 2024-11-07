package main

import (
	"time"

	"github.com/sgs921107/glogging"
	"go.uber.org/zap"
)

func main() {
	logging := glogging.NewZapLogging(glogging.Options{
		Level:          "info",
		FilePath:       "./log.log",
		RotationTime:   time.Minute,
		RotationMaxAge: time.Minute * 3,
		Caller:         "full",
		UseUTC:         true,
		TimeFormater:   "2006-01-02 15:04:05-07:00",
	})
	defer logging.Sync()
	logger := logging.GetLogger()
	sugaredLogger := logging.GetSugaredLogger()
	logger.Info("hello world!")
	sugaredLogger.Info("hello")
	sugaredLogger.Infow("hello world",
		"logger", "zap sugaredLogger")
	zap.L().Info("hello")
	zap.S().Info("hello")
}
