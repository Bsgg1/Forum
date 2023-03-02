package main

import (
	"forum/common"
	"forum/dao/mysql"
	"forum/routers"
)

func main() {
	err := common.InitViper()
	if err != nil {
		panic(err)
	}
	mysql.InitMysql()
	mysql.GenModel()
	r := routers.Router()
	r.Run("localhost:8080")
}
