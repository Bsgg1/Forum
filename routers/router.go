package routers

import (
	"forum/logic/post"
	"forum/logic/user"
	"forum/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	r.POST("/user/register", user.Register)
	r.POST("/auth", user.Login)
	r.POST("/user/login", user.Login)
	r.POST("/user/dolike", middleware.JWTAuthMiddleware(), user.DoLike)
	r.POST("/user/createpost", middleware.JWTAuthMiddleware(), post.CreatePost)
	r.POST("/user/postlist", post.List)
	return r
}
