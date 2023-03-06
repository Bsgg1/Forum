package user

import (
	"forum/common"
	"forum/dao/mysql"
	"forum/middleware"
	"forum/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req model.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  "无效的参数",
		})
		return
	}
	user := mysql.FindByName(req.UserName)
	psd := common.MD5(req.PassWord)
	if user.UserName == "" || psd != user.PassWord {
		c.JSON(-1, common.Message{
			Code: -1,
			Msg:  "用户名或者密码错误",
		})
		return
	}
	tokenString, _ := middleware.GenToken(user.UserName, user.ID)
	c.JSON(200, common.Message{
		Code: 0,
		Msg:  tokenString,
	})
}
