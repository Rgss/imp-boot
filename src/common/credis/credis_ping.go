package credis

import (
	"github.com/robfig/cron"
	"log"
	"time"
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


func ping () {
	log.Printf("[info] redis ping ...")
	for key, value := range connection.pool  {
		_, err := value.Ping().Result()
		if err != nil {
			log.Printf("[error] redis#%v ping error: %v", key, err.Error())
		}
	}
}