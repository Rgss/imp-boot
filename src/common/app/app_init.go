package app

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"github.com/Rgss/imp-boot/src/common/clog"
	"github.com/Rgss/imp-boot/src/common/config"
	"github.com/Rgss/imp-boot/src/common/credis"
	"github.com/Rgss/imp-boot/src/common/db"
)

// app
var application *Application

// base path
var basePath string

// global app config
//var appConfig *config.AppConfig

// logger
//var appLogger *clog.CLogger

// init app
func InitApplication() {
	if application != nil && application.inited {
		return
	}

	log.Printf("[info] InitApplication...")
	initCliParams()

	config.Initialize(confPath)
	//log.Printf("[info] configs: %v", configs)

	db.Initialize(Configure().GetDatabaseConfig())

	application = newApplication(config.Appconfig())

	clog.Initialize(application.GetLogPath())

	credis.Initialize(Configure().GetRedisConfig())

	//initPid()
}

// 检测程序是否已经执行
func initPid() {
	pidPath := application.GetPidPath()
	var err = tryPid(pidPath)
	if err != nil {
		log.Println(err)
		os.Exit(100)
	}
}


// app
func Core() *Application {
	log.Printf("[info] call app.Core ... ")
	if application == nil || !application.inited {
		InitApplication()
		//debug.PrintStack()
		//panic("app is not inited.")
	}
	return application
}

// config
func Configure() *config.AppConfig {
	return config.Appconfig()
}

// logger
func Logger() *logrus.Logger {
	return clog.Clogger().GetLogger()
}

func Redisx() *credis.CRedis {
	return Cache()
}

func Cache() *credis.CRedis {
	return credis.Client()
}