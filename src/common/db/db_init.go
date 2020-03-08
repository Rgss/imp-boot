package db

import (
	"github.com/jinzhu/gorm"
	"log"
	"github.com/Rgss/imp-boot/src/common/config"
)

var connection *DB

func Initialize(configs map[string]*config.DatabaseConfig) {

	// 实例化
	NewDB(configs)

	// ping
	go Ping()
}


// new DB
func NewDB(configs map[string]*config.DatabaseConfig) *DB {

	connection = &DB{
		DefaultName: "default",
	}
	pool := make(map[string]*gorm.DB, 0)

	for key, config := range configs {

		url 	:= connection.getUrl(config)
		dialect := config.Driver
		maxIdleConns := config.MaxIdleConns
		maxOpenConns := config.MaxOpenConns
		enableLog 	 := config.EnableLog

		db, err := gorm.Open(dialect, url);
		if err != nil {
			log.Printf("连接数据库失败 error: %s", err.Error())
			//app.Logger().Errorf("连接数据库失败 error: %s", err.Error())
			panic("连接数据库失败 error: " + err.Error())
			return nil
		}

		db.LogMode(enableLog)
		db.SingularTable(true) // 禁用表名负数
		db.DB().SetMaxIdleConns(maxIdleConns)
		db.DB().SetMaxOpenConns(maxOpenConns)

		//if err = db.AutoMigrate(models...).Error; nil != err {
		//	log.Errorf("auto migrate tables failed: %s", err.Error())
		//}
		//return

		db.Debug()

		mapID := connection.DBID(key)
		pool[mapID] = db
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "mnt_" + defaultTableName
	//}

	connection.pool = pool
	return connection
}


// 对外暴露
func Connection() (*DB) {
	return connection
}