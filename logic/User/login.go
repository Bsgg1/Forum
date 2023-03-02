package User

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
		c.JSON(-1, common.Message{
			Code: 0,
			Msg:  err.Error(),
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
	j := middleware.NewJWT()
	token, err := j.GenToken(user.UserName, user.ID)
	if err != nil {
		c.JSON(-1, common.Message{
			Code: 0,
			Msg:  "获取token失败",
		})
	}
	c.JSON(0, common.Message{
		Code: 1,
		Msg:  token,
	})
	parseToken, _ := j.ParseToken(token)
	c.JSON(0, parseToken)
}
