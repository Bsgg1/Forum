package main

import (
	"forum/common"
	"forum/dao/mysql"
	"forum/dao/redis"
	"forum/routers"
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
	r.Run("localhost:8080")
}
