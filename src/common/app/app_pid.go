package app

import (
	"errors"
	"os"
	"strconv"
	"syscall"
	"log"
)

// create pid file
func tryPid(pidFilePath string) error {
	_, err := os.Stat(pidFilePath)
	if err == nil {
		err = errors.New("[error] the pid file exists, please delete this file before startting the server. path: " + pidFilePath)
		return err
	}

	//return nil;
	name := pidFilePath
	file, err := os.OpenFile(name, syscall.O_CREAT|syscall.O_RDWR|syscall.O_CLOEXEC, 0777)
	if err != nil {
		return err
	}

	defer file.Close()

	var pid int
	pid = os.Getpid() // int
	var pidStr = strconv.Itoa(pid)
	_, err = file.WriteString(pidStr)

	return err
}

// remove pid file
func removePid(pidFilePath string) error {
	err := os.Remove(pidFilePath)
	if err != nil {
		log.Printf("[error] remove the pid file fail. %v", err.Error())
		return err
	}
	return nil;
}