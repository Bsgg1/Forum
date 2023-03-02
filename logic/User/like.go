package User

import (
	"forum/common"
	"forum/dao/redis"
	"forum/middleware"
	"forum/model"
	"github.com/gin-gonic/gin"
	"time"
)

func DoLike(c *gin.Context) {
	var info model.UserDoLike
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(-1, common.Message{
			Code: 0,
			Msg:  "error",
		})
		return
	}
	c.JSON(0, info.Token)
	j := middleware.NewJWT()
	claims, err := j.ParseToken(info.Token)
	if err != nil {
		c.JSON(-1, common.Message{
			Code: 0,
			Msg:  "error",
		})
		return
	}
	if info.Status == "like" {
		key := "like--" + claims.UserName + "--" + info.ToName
		_, err := redis.Rank(key, info.ToName)
		if err == common.RepeatedFollow {
			c.JSON(-1, common.Message{
				Code: 0,
				Msg:  common.RepeatedFollow.Error(),
			})
			return
		}
		c.JSON(1, key)
		err = redis.AddLike(key, info.ToName, float64(time.Now().Unix()))
		if err != nil {
			c.JSON(0, err.Error())
			return
		}
		c.JSON(0, "success")
	} else {

	}

}
