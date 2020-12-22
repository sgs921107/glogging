package main

import (
	"time"

	"github.com/sgs921107/glogging"
)

func main() {
	logging := glogging.NewLogging(&glogging.Options{
		Level: "debug",
		FilePath: "./log.log",
		RotationTime: time.Minute,
		RotationMaxAge: time.Minute * 3,
	})
	logger := logging.GetLogger()
	logger.Info("hello")
}