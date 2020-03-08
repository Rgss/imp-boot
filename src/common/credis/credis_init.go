package credis

import (
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"log"
	"github.com/Rgss/imp-boot/src/common/config"
)

var connection *CRedis
var redisConfigs map[string]*config.RedisConfig

// 初始化redis
func Initialize(config map[string]*config.RedisConfig) {
	//log.Printf("initRedis ...")
	redisConfigs = config
	connection = NewCRedisClient("default")

	go Ping()
}


// new client
func NewCRedisClient(name string) *CRedis {
	config, err := GetConfig(name)
	if err != nil {
		panic("the redis config " + name + " not exists!")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.Db,       // use default DB
	})

	connection = &CRedis{
		Client: client,
		Config: config,
	}

	var pools = make(map[string] *redis.Client)
	nameId := connection.RedisID(name)
	pools[nameId] = client
	connection.pool = pools

	//log.Printf("redisClient: %v", redis)

	return connection
}

// config
func GetConfig(name string) (*config.RedisConfig, error) {
	var config = &config.RedisConfig{}
	if _, exists := redisConfigs[name]; exists {
		config = redisConfigs[name]
	} else {
		log.Printf("the redis config %v not exists", name)
		//log.Printf("redis config %v %v %v", name, config, appConfig.Redis)
		return config, errors.New("the redis config " + name + " not exists")
	}

	return config, nil
}

// redis
func Client() *CRedis {
	return connection
}
