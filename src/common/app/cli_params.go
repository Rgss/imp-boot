package app

import (
	"flag"
	"log"
	"github.com/Rgss/imp-boot/src/common/config"
)

// 命令行参数
var cliParams map[string]string

// 项目配置文件路径
var confPath = *flag.String("c", "", "the app config path.")

// init cli params
func initCliParams() {
	flag.Parse()

	if len(confPath) <= 0 {
		confPath = config.GetDefaultConfPath()
		log.Printf("[notice] no confPath. use the default confPath: %v", confPath)
	}
}
