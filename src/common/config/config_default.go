package config

import (
	"github.com/Rgss/imp-boot/src/common/utils"
)


// 默认pid路径
func GetDefaultPidPath() (string) {
	var pidPath = utils.GetBasePath() + "monitor.pid";
	return pidPath
}

// 默认log路径
func GetDefaultLogPath() (string) {
	var logPath = utils.GetBasePath() + "logs/";
	return logPath
}

// 默认任务文件路径
func GetDefaultTaskSavePath() (string) {
	var taskSavePath = utils.GetBasePath() + "conf/job.yaml";
	return taskSavePath
}

// 默认任务缓存路径
func GetDefaultTaskCachePath() (string) {
	var taskSavePath = utils.GetBasePath() + "conf/cache.txt";
	return taskSavePath
}