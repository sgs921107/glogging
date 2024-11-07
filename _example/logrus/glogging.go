package main

import (
	"time"

	"github.com/sgs921107/glogging"
)

func main() {
	logging := glogging.NewLogrusLogging(glogging.Options{
		Level:          "info",
		FilePath:       "./log.log",
		RotationTime:   time.Minute,
		RotationMaxAge: time.Minute * 3,
		Caller:         "full",
		UseUTC:         true,
		TimeFormater:   "2006-01-02 15:04:05-07:00",
	})
	logger := logging.GetLogger()
	for i := 1; i < 5; i++ {
		logger.WithFields(glogging.LogrusFields{
			"logger": "logrus",
		}).Info("hello world!")
	}
}
