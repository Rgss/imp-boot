package db

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Rgss/imp-boot/src/common/utils"
	"github.com/Rgss/imp-boot/src/common/config"
)

type DB struct {
	DefaultName string
	pool        map[string]*gorm.DB
}

// url
func (mc *DB) getUrl(config *config.DatabaseConfig) string {
	url := fmt.Sprintf("%s:%s@/%s?charset=%s", config.User, config.Pass, config.DbName, config.Charset)
	return url
}

// gorm db
func (mc *DB) DB(args... string) *gorm.DB {
	dbName := mc.DefaultName
	for _, val := range args {
		dbName = val
	}

	mapID := mc.DBID(dbName)
	if _, ok := mc.pool[mapID]; !ok {
		log.Printf("[error] db %v is not exists", dbName)
	}

	return mc.pool[mapID]
}

// hash id
func (mc *DB) DBID(dbName string) string {
	dbName = "db_" + utils.Md5(dbName)
	return dbName
}
