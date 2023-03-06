package post

import (
	"forum/common"
	"forum/dao/mysql"
	"forum/dao/redis"
	"forum/model"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	userid := c.MustGet("userid").(uint)
	username := c.MustGet("username").(string)
	postinfo := &model.Post{
		UserId:    userid,
		Content:   c.PostForm("content"),
		LikeCount: 0,
	}
	if err := mysql.NewPost(postinfo); err != nil {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	_, err := redis.DelPostsList(username + "--posts")
	if err != nil {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, common.Message{
		Code: 0,
		Msg:  "success",
	})
}
