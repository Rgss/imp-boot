package config

import (
	"errors"
)

type AppConfig struct {
	Base     *BaseConfig                `yaml:"base"`
	Http     *HttpConfig                `yaml:"http"`
	Redis    map[string]*RedisConfig    `yaml:"redis"`
	Database map[string]*DatabaseConfig `yaml:"database"`
	App      map[string]interface{}
}

// base config
type BaseConfig struct {
	BasePath      string `yaml:"basePath"`
	PidPath       string `yaml:"pidPath"`
	LogPath		  string `yaml:"logPath"`
	TaskSavePath  string `yaml:"taskSavePath"`
	TaskCachePath string `yaml:"taskCachePath"`
}

// http config
type HttpConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type EmailConfig struct {
	Identity  string
	Username  string
	Password  string
	Host      string
	Port      int
	Aliasname string
}

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Driver       string `yaml:"driver"`
	User         string `yaml:"user"`
	Pass         string `yaml:"pass"`
	Charset      string `yaml:"charset"`
	DbName       string `yaml:"dbName"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	EnableLog    bool   `yaml:"enableLog"`
	TablePrefix  string `yaml:"tablePrefix"`
}

// get app system config. as base redis http
//func (appConf *AppConfig) GetSystem(name string, module string) (interface{}, error) {
//	appConfigs := make(map[string]interface{}, 0)
//	switch strings.ToLower(module) {
//	case "base":
//		appConfigs = appConfig.Base
//	case "http":
//		appConfigs = appConfig.Http
//	}
//
//	var config interface{}
//	if _, exists := appConfigs[name]; exists {
//		config = appConfigs[name]
//	} else {
//		log.Printf("the system " + module + " config " + name + " not exists")
//		return nil, errors.New("the system " + module + " config " + name + " not exists")
//	}
//
//	return config, nil
//}

// get app config
func (appConf *AppConfig) Get(name string) (interface{}, error) {
	var config interface{}
	if _, exists := appConf.App[name]; exists {
		config = appConf.App[name]
	} else {
		return nil, errors.New("the redis config " + name + " not exists")
	}

	return config, nil
}

// set app config
func (appConf *AppConfig) Set(name string, value interface{}) (bool, error) {
	appConf.App[name] = value
	return true, nil
}

func (appConf *AppConfig) GetRedisConfig() (map[string]*RedisConfig) {
	return appConf.Redis
}

func (appConf *AppConfig) GetHttpConfig() (*HttpConfig) {
	return appConf.Http
}

func (appConf *AppConfig) GetDatabaseConfig() (map[string]*DatabaseConfig) {
	return appConf.Database
}
