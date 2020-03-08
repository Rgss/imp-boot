package credis

import (
	"github.com/go-redis/redis/v7"
	"log"
	"github.com/Rgss/imp-boot/src/common/config"
	"github.com/Rgss/imp-boot/src/common/utils"
)

const(
	DefaultName = "default"
)

type CRedis struct {
	DefaultName string
	Config *config.RedisConfig
	Client *redis.Client
	pool   map[string]*redis.Client
}

type RedisNode struct {
	RedisConfigs map[string] config.RedisConfig
}

type connect struct {
	config config.RedisConfig `yaml:"redis"`
}

//
func (this *CRedis) DB(name... string) *redis.Client {
	dbName := DefaultName
	for _, val := range name {
		dbName = val
	}

	mapID := this.RedisID(dbName)
	if _, ok := this.pool[mapID]; !ok {
		log.Printf("[error] redis %v is not exists", dbName)
	}

	return this.pool[mapID]
}

func (this *CRedis) RedisID(name string) string {
	name = "redis_" + utils.Md5(name)
	return name
}


