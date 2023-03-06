package user

import (
	"forum/common"
	"forum/dao/mysql"
	"forum/dao/redis"
	"github.com/gin-gonic/gin"
)

func DoLike(c *gin.Context) {
	username := c.MustGet("username").(string)
	toname := c.PostForm("toname")
	status := c.PostForm("status")
	var st = 1
	if status == "unlike" {
		st = -1
	}
	c.JSON(200, st)
	f1, err := redis.FindRelation("user-relation", "like--"+toname+"--"+username)
	if err != nil {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	f2, err := redis.FindRelation("user-relation", "unlike--"+toname+"--"+username)
	if err != nil {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	cnt := 0
	if st == 1 {
		if f1 {
			c.JSON(200, common.Message{
				Code: 0,
				Msg:  common.RepeatedFollow.Error(),
			})
			return
		} else if f2 {
			_, err = redis.AddRelation("user-relation", "like--"+toname+"--"+username)
			if err != nil {
				c.JSON(200, common.Message{
					Code: 0,
					Msg:  err.Error(),
				})
				return
			}
			_, err = redis.DelRelation("user-relation", "unlike--"+toname+"--"+username)
			if err != nil {
				c.JSON(200, common.Message{
					Code: 0,
					Msg:  err.Error(),
				})
				return
			}
			cnt = 2
		} else {
			_, err = redis.AddRelation("user-relation", "like--"+toname+"--"+username)
			if err != nil {
				c.JSON(200, common.Message{
					Code: 0,
					Msg:  err.Error(),
				})
				return
			}
			cnt = 1
		}
	} else {
		if f2 {
			c.JSON(200, common.Message{
				Code: 0,
				Msg:  common.RepeatedFollow.Error(),
			})
			return
		} else if f1 {
			_, err = redis.AddRelation("user-relation", "unlike--"+toname+"--"+username)
			if err != nil {
				c.JSON(200, common.Message{
					Code: 0,
					Msg:  err.Error(),
				})
				return
			}
			_, err = redis.DelRelation("user-relation", "like--"+toname+"--"+username)
			if err != nil {
				c.JSON(200, common.Message{
					Code: 0,
					Msg:  err.Error(),
				})
				return
			}
			cnt = -2
		} else {
			_, err = redis.AddRelation("user-relation", "unlike--"+toname+"--"+username)
			if err != nil {
				c.JSON(200, common.Message{
					Code: 0,
					Msg:  err.Error(),
				})
				return
			}
			cnt = -1
		}
	}
	if err := mysql.UpdateRelation(username, toname, st); err != nil {
		c.JSON(200, common.Message{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	to := mysql.FindByName(toname)
	to.FollowerCount += cnt
	mysql.UpdateUserInfo(to)
	user := mysql.FindByName(username)
	user.LikeCount += st
	mysql.UpdateUserInfo(user)
	c.JSON(200, common.Message{
		0,
		"success",
	})
}
