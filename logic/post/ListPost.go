package post

import (
	"encoding/json"
	"forum/common"
	"forum/dao/mysql"
	"forum/dao/redis"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	username := c.PostForm("username")
	if username == "" {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  "用户名不符合规则",
		})
		return
	}
	count, err := redis.FindUser(username)
	if err != common.RedisNil {
		c.JSON(200, common.Message{
			Code: -100,
			Msg:  err.Error(),
		})
		return
	}
	if count != 0 {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  "用户名不存在",
		})
		return
	}
	user := mysql.FindByName(username)
	if user.UserName == "" {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  "用户名不存在",
		})
		if err := redis.AddNotExistUser(username); err != nil {
			c.JSON(200, common.Message{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		return
	}
	list, err := redis.FindPostsList(username + "--posts")
	if err != common.RedisNil {
		c.JSON(200, common.Message{
			Code: 0,
			Msg:  list,
		})
		return
	}
	posts, err := mysql.ListPost(user.ID)
	if err != nil {
		c.JSON(200, common.Message{
			Code: -100,
			Msg:  err.Error(),
		})
		return
	}
	p, err := json.Marshal(posts)
	if err != nil {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	redis.AddPostsList(username+"--posts", string(p))

	c.JSON(200, common.Message{
		Code: 0,
		Msg:  string(p),
	})
}
