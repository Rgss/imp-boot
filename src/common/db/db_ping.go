package db

import (
	"time"
	"github.com/robfig/cron"
	"log"
)

func Ping() {
	loc, _  := time.LoadLocation("PRC")
	dbcron  := cron.NewWithLocation(loc)

	// 检查心跳
	dbcron.AddFunc("@every 30s", func() {
		ping()
	})

	dbcron.Start()
}

func ping() {
	log.Printf("[info] database ping ...")
	for key, value := range connection.pool  {
		err := value.DB().Ping()
		if err != nil {
			log.Printf("[error] db#%v ping error: %v", key, err.Error())
		}
	}
}