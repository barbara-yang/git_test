package redis

import (
	"entry_task/common/config"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var redisCli *Cli

// Cli is a Singleton Object
type Cli struct {
	*redis.Client
}

// InitRedis is creat a redis-cli pool for user
func InitRedis() {
	newRedisCli := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    fmt.Sprintf("%s:%d", config.REDISSREVER, config.REDISPORT),

		PoolSize:     config.REDISMAXOPEN,
		MinIdleConns: config.REDISMAXIDLE,
		ReadTimeout:  config.REDISTIMEOUT, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: config.REDISTIMEOUT, //写超时，默认等于读超时
		MaxRetries:   config.REDISRETRY,
	})
	redisCli = &Cli{newRedisCli}

}

// GetRedisCli return the singleton object to user
func GetRedisCli() *Cli {
	if redisCli == nil {
		log.Fatal("Mysql init error.")
	}
	return redisCli
}
