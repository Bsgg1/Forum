package routers

import (
	"forum/logic/User"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	r.POST("/user/register", User.Register)
	r.POST("/user/login", User.Login)
	r.POST("/user/dolike", User.DoLike)
	return r
}
