package glogging

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type baseLog struct {
	options	Options
}

// cratePattern create filename pattern
func (l *baseLog) createPattern() string {
	var p string
	duration := l.options.RotationTime
	switch {
	case duration < time.Hour:
		p = ".%Y%m%d%H%M"
	case duration < time.Hour * 24:
		p = ".%Y%m%d%H"
	case duration >= time.Hour * 24:
		p = ".%Y%m%d"
	default:
		p = ".%Y%m%d%H"
	}
	return p
}

// log output
func (l *baseLog) output() io.Writer {
	filePath := l.options.FilePath
	if filePath == "" {
		return os.Stdout
	}
	if l.options.RotationMaxAge != 0 && l.options.RotationMaxAge < l.options.RotationTime {
		l.options.RotationMaxAge = l.options.RotationTime
	}
	if writer, err := rotatelogs.New(
		filePath+l.createPattern(),
		rotatelogs.WithLinkName(filePath),
		rotatelogs.WithRotationTime(l.options.RotationTime),
		rotatelogs.WithMaxAge(l.options.RotationMaxAge),
		// rotatelogs.ForceNewFile(),
	); err != nil {
		panic("Set Log File Failed: " + err.Error())
	} else {
		return writer
	}
}