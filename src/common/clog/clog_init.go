package clog

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// clogger
var clogger *CLogger

// init
func Initialize(logPath string) {
	appLogger := logrus.New()

	clogger = &CLogger{
		logger:  appLogger,
		logPath: logPath,
	}

	logFile := clogger.GetLogFile()
	//log.Printf("file: %v", logFile)
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		clogger.logger.SetOutput(file)
	} else {
		log.Printf("[warn] appLogger SetOutput fail, using default stderr")
	}

	//appLogger.Info("appLogger logrus ....")
}

func Clogger() (*CLogger) {
	return clogger
}
