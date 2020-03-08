package app

import (
	"github.com/Rgss/imp-boot/src/common/config"
	"github.com/Rgss/imp-boot/src/common/utils"
)

type Application struct {
	inited bool
	pid    string
	config *config.AppConfig
}

// new Application
func newApplication(config *config.AppConfig) (*Application) {
	application := &Application{
		inited: true,
		config: config,
	}

	return application
}

// GetHttpConfig
func (core *Application) GetHttpConfig() (*config.HttpConfig) {
	return core.config.Http;
}

// GetBasePath
func (core *Application) GetBasePath() (string) {
	return core.config.Base.BasePath
}

// GetPidPath
func (core *Application) GetPidPath() (string) {
	return core.config.Base.PidPath
}

// GetLogPath
func (core *Application) GetLogPath() (string) {
	return core.config.Base.LogPath
}

// GetTaskSavePath
func (core *Application) GetTaskSavePath() (string) {
	return core.config.Base.TaskSavePath
}

// GetTaskCachePath
func (core *Application) GetTaskCachePath() (string) {
	return core.config.Base.TaskCachePath
}

// get Application config
func (core *Application) GetConfig(name string) (interface{}, error) {
	config, err := core.config.Get(name);
	return config, err
}

// destroy resource when the program stop
func (core *Application) Destroy() {
	removePid(core.config.Base.PidPath)
}


func (core *Application) IsInited() (bool) {
	return core.inited
}

// 基础目录
func GetBasePath() (string) {
	return utils.GetBasePath()
}

// 获取当前路径
func GetCurrentPath() (string, error) {
	return utils.GetCurrentPath()
}

