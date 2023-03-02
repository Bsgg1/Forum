package model

import (
	"encoding/hex"
	"encoding/json"
	"gorm.io/gorm"
)

type UserInfo struct {
	*gorm.Model
	UserName      string `gorm:"unique;not null;type:varchar(64)"`
	PassWord      string `gorm:"not null;type:varchar(64)"`
	Token         string `gorm:"type:varchar(64)"`
	Hash          string `gorm:"not null;type:varchar(64)"`
	FollowerCount int
	LikeCount     int
}
type UserRelation struct {
	*gorm.Model
	User1  string
	User2  string
	status int8 //1关注 0未关注
}

func (UserRelation) TableName() string {
	return "user_relation"
}

func (user *UserInfo) GetUid() string {
	bytes, err := json.Marshal(map[string]interface{}{
		"userid":   user.ID,
		"username": user.UserName,
	})
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

type UserRegister struct {
	UserName   string `json:"username" binding:"required"`
	PassWord   string `json:"password" binding:"required"`
	RePassWord string `json:"repassword" binding:"required"`
}
