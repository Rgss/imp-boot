package utils

import (
	"os/exec"
	"os"
	"path/filepath"
	"strings"
	"errors"
)

var basePath string

// 基础目录
func GetBasePath() (string) {
	if len(basePath) > 0 {
		return basePath
	}

	basePath, _ = GetCurrentPath()
	return basePath
}

// 获取当前路径
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
