package main

import (
	"time"

	"go.uber.org/zap"
	"github.com/sgs921107/glogging"
)

func main() {
	logging := glogging.NewZapLogging(glogging.Options{
		Level: "info",
		FilePath: "./log.log",
		RotationTime: time.Minute,
		RotationMaxAge: time.Minute * 3,
	})
	defer logging.Sync()
	logger := logging.GetLogger()
	sugaredLogger := logging.GetSugaredLogger()
	logger.Info("hello")
	sugaredLogger.Info("hello")
	zap.L().Info("hello")
	zap.S().Info("hello")
}