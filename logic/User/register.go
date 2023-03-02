package User

import (
	"errors"
	"forum/common"
	"forum/dao/mysql"
	"forum/model"
	"github.com/gin-gonic/gin"
	"time"
)

var Info model.UserRegister

func Register(c *gin.Context) {
	if err := c.ShouldBindJSON(&Info); err != nil {
		c.JSON(-1, common.Message{
			Code: 0,
			Msg:  "error",
		})
		return
	}
	if Info.PassWord != Info.RePassWord {
		c.JSON(-1, common.Message{
			Code: 0,
			Msg:  "error",
		})
		return
	}
	user := mysql.FindByName(Info.UserName)
	if user.UserName != "" {
		c.JSON(-1, common.Message{
			Code: -1,
			Msg:  errors.New("该用户已经存在").Error(),
		})
		return
	}
	UserInfo := model.UserInfo{
		UserName:      Info.UserName,
		PassWord:      common.MD5(Info.PassWord),
		Hash:          time.Now().String(),
		FollowerCount: 0,
		LikeCount:     0,
	}
	if err := mysql.CreateUser(&UserInfo); err != nil {
		c.JSON(0, common.Message{
			Code: 0,
			Msg:  "error",
		})
		return
	}
	c.JSON(0, common.Message{
		Code: 1,
		Msg:  "success",
	})
}
