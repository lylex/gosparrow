package log

import (
	"fmt"

	"github.com/Sirupsen/logrus"

	"gosparrow/pkg/consts"
	"gosparrow/pkg/utils"
)

// InitLog used to initialize a logger with log level, default is info level
func InitLog(level string) error {
	if level != "" {
		var logLevel logrus.Level
		var err error
		if logLevel, err = logrus.ParseLevel(level); err != nil {
			return fmt.Errorf(consts.ErrLogLevelParsePattern, level)
		}
		logrus.SetLevel(logLevel)
	}

	return nil
}

// App used to write down app level logs
func App(level, msg string) {
	logFormat := "%s [APP] %s"
	now := utils.Now()
	args := []interface{}{
		now,
		msg,
	}

	perform(level, logFormat, args...)
}

// Individual used to write down Individual level logs
func Individual(level, requestID, requestIP, method, uri, msg string) {
	logFormat := "%s [%s] %s %s %s %s"
	now := utils.Now()
	args := []interface{}{
		now,
		requestID,
		requestIP,
		method,
		uri,
		msg,
	}

	perform(level, logFormat, args...)
}

func perform(level, logFormat string, args ...interface{}) {
	switch level {
	case "panic":
		logrus.Panicf(logFormat, args...)
	case "fatal":
		logrus.Fatalf(logFormat, args...)
	case "error":
		logrus.Errorf(logFormat, args...)
	case "warn":
		logrus.Warnf(logFormat, args...)
	case "info":
		logrus.Infof(logFormat, args...)
	case "debug":
		logrus.Debugf(logFormat, args...)
	default:
		logrus.Panicf("Unknown log level %s", level)
	}
}
