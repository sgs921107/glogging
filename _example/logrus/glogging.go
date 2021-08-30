package main

import (
	"time"

	"github.com/sgs921107/glogging"
)

func main() {
	logging := glogging.NewLogrusLogging(glogging.Options{
		Level: "info",
		FilePath: "./log.log",
		RotationTime: time.Minute,
		RotationMaxAge: time.Minute * 3,
	})
	logger := logging.GetLogger()
	logger.Info("hello")
}