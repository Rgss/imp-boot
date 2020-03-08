package clog

import (
	"github.com/sirupsen/logrus"
	"time"
)


type CLogger struct {
	logPath string
	logger  *logrus.Logger
}

// get log file
func (clogger *CLogger) GetLogFile() (string) {
	d := time.Now().Format("2006-01-02")
	var logPath = clogger.logPath + d + ".log"
	return logPath
}

// logrus.Logger
func (clogger *CLogger) GetLogger() (*logrus.Logger) {
	return clogger.logger
}

// log path
func (clogger *CLogger) GetLogPath() (string) {
	return clogger.logPath
}
