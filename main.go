package main

import (
	"fmt"
	"forum/common"
	"forum/dao/mysql"
	"forum/dao/redis"
	"forum/routers"
	"github.com/spf13/viper"
)

func main() {
	err := common.InitViper()
	if err != nil {
		panic(err)
	}
	mysql.InitMysql()
	mysql.GenModel()
	redis.InitRedis()
	r := routers.Router()
	fmt.Println(viper.GetString("Redis.Addr"),
		viper.GetString("Redis.Password"),
		viper.GetInt("Redis.DB"),
		viper.GetInt("Redis.PoolSize"),
		viper.GetInt("Redis.MinIdleConn"))
	r.Run("localhost:8080")
}
