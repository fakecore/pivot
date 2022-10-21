package core

import (
	"github.com/go-redis/redis/v9"
	"povit/global"
)

func InitRedisOrDie() *redis.Client {

	redis := redis.NewClient(&redis.Options{
		Addr:     global.G_CONF.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//FIX ME check whether redis connected

	return redis

}
