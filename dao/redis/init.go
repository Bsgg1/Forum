package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RD *redis.Client

func InitRedis() {
	RD = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("Redis.Addr"),
		Password:     viper.GetString("Redis.Password"),
		DB:           viper.GetInt("Redis.DB"),
		PoolSize:     viper.GetInt("Redis.PoolSize"),
		MinIdleConns: viper.GetInt("Redis.MinIdleConn"),
	})

}
