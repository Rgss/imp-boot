package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

var appConfig *AppConfig

// init app config
func Initialize(confPath string) (*AppConfig) {
	appConfig = &AppConfig{}

	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalf("[error] open the config file error, please check the file %v.\n %v", confPath, err)
		return appConfig
	}

	err = yaml.Unmarshal(yamlFile, &appConfig)
	if err != nil {
		log.Fatalf("[error] the config file Unmarshal error, please check the file %v.\n %v", confPath, err)
	}

	//log.Printf("confPath: %v", *confPath)
	//log.Printf("appConfig: %v", appConfig.Base)

	if len(appConfig.Base.PidPath) <= 0 {
		appConfig.Base.PidPath = GetDefaultPidPath()
	}

	if len(appConfig.Base.LogPath) <= 0 {
		appConfig.Base.LogPath = GetDefaultLogPath()
	}

	if len(appConfig.Base.TaskSavePath) <= 0 {
		appConfig.Base.TaskSavePath = GetDefaultTaskSavePath()
	}

	if len(appConfig.Base.TaskCachePath) <= 0 {
		appConfig.Base.TaskCachePath = GetDefaultTaskCachePath()
	}

	return appConfig
}

func Appconfig() (*AppConfig) {
	return appConfig
}